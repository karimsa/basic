// cmd/asm/asm.go
// Tiny assembler.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/karimsa/basic/debug"
	"github.com/karimsa/basic/program"
)

type unwrittenInstruct struct {
	addr uint16
	cmd  string
	args []string
}

var (
	simpleInstructions = map[string]uint16{
		// RRI full
		"CLA": 0x7800,
		"CLE": 0x7400,
		"CMA": 0x7200,
		"CME": 0x7100,
		"CIR": 0x7080,
		"CIL": 0x7040,
		"INC": 0x7020,
		"SPA": 0x7010,
		"SNA": 0x7008,
		"SZA": 0x7004,
		"SZE": 0x7002,
		"HLT": 0x7001,

		// IOI full
		"INP": 0xF800,
		"UT":  0xF400,
		"SKI": 0xF200,
		"SKO": 0xF100,
		"ION": 0xF080,
		"IOP": 0xF040,
	}

	// left ops are direct, right are indirect
	mriInstructions = map[string][2]uint16{
		"AND": [2]uint16{0x0000, 0x8000},
		"ADD": [2]uint16{0x1000, 0x9000},
		"LDA": [2]uint16{0x2000, 0xA000},
		"STA": [2]uint16{0x3000, 0xB000},
		"BUN": [2]uint16{0x4000, 0xC000},
		"BSA": [2]uint16{0x5000, 0xD000},
		"ISZ": [2]uint16{0x6000, 0xE000},
	}
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("usage: asm <program> [outfile]")
		return
	}

	input := os.Args[1]

	output := "./a.out"
	if len(os.Args) == 3 {
		output = os.Args[2]
	}

	if debug.ASM {
		fmt.Printf("Will compile to: %s\n", output)
	}

	inputFd, err := os.Open(input)
	if err != nil {
		fmt.Printf("Failed to open program: %s\n", err.Error())
		return
	}

	w, err := program.WriteProgram(output)
	if err != nil {
		fmt.Printf("Failed to create output: %s\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(inputFd)
	addr := uint16(0)
	unwritten := make([]*unwrittenInstruct, 0)

	for scanner.Scan() {
		line := scanner.Text()
		commentLoc := strings.Index(line, ";")
		if commentLoc != -1 {
			line = line[0:commentLoc]
		}
		line = strings.Trim(line, " ")
		if line == "" {
			continue
		}

		label := ""
		instructLine := line

		if labelSplit := strings.Split(line, ":"); len(labelSplit) != 1 {
			label = labelSplit[0]
			instructLine = strings.Trim(labelSplit[1], " ")

			if len(labelSplit) > 2 {
				fmt.Printf("too many labels: %#v\n", labelSplit)
				return
			}
		}
		instruct := strings.Split(instructLine, " ")
		if debug.ASM {
			fmt.Printf("Read: [%s] %s %s\n", label, instruct[0], instruct[1:])
		}

		if strCmp(instruct[0], "org") {
			if len(instruct) != 2 {
				fmt.Printf("Syntax-error: expected hex, then EOL: %s\n", strings.Join(instruct, " "))
				return
			}

			org, err := strconv.ParseInt(instruct[1], 16, 16)
			if err != nil {
				fmt.Printf("Non-hexadecimal argument given to ORG: %s\n", instruct[1])
				return
			}

			addr = uint16(org)
		} else {
			// set addr of label
			if label != "" {
				w.SetLabel(label, addr)
			}

			// direct write any raw data
			if strCmp(instruct[0], "hex") {
				word, err := strconv.ParseInt(instruct[1], 16, 16)
				if err != nil {
					fmt.Printf("Non-hexadecimal argument given to HEX: %s\n", instruct[1])
					return
				}

				w.SetWord(addr, uint16(word))
			} else if strCmp(instruct[0], "dec") {
				word, err := strconv.ParseInt(instruct[1], 10, 16)
				if err != nil {
					fmt.Printf("Non-decimal argument given to DEC: %s\n", instruct[1])
					return
				}

				w.SetWord(addr, uint16(word))
			} else if instWord, ok := simpleInstructions[instruct[0]]; ok {
				if len(instruct) != 1 {
					fmt.Printf("Syntax error: %s does not take arguments\n", instruct[0])
					return
				}

				w.SetWord(addr, instWord)
			} else {
				unwritten = append(unwritten, &unwrittenInstruct{
					addr: addr,
					cmd:  instruct[0],
					args: instruct[1:],
				})
			}
		}

		if !strCmp(instruct[0], "org") {
			addr++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read program: %s\n", err.Error())
		return
	}

	for _, instruct := range unwritten {
		if len(instruct.args) == 0 || len(instruct.args) > 2 {
			fmt.Printf("Syntax error: unexpected set of args given to %s (takes 1 or 2): %#v\n", instruct.cmd, instruct.args)
			return
		}

		mriInstruct, ok := mriInstructions[instruct.cmd]
		if !ok {
			fmt.Printf("Syntax error: unknown instruction: %s\n", instruct.cmd)
			return
		}

		opcode := mriInstruct[0]
		if len(instruct.args) == 2 {
			if len(instruct.args[1]) != 1 || unicode.ToUpper(rune(instruct.args[1][0])) != 'I' {
				fmt.Printf("Syntax error: second arg to an MRI can only be I\n")
				return
			}

			opcode = mriInstruct[1]
		}

		labelAddr, ok := w.GetLabel(instruct.args[0])
		if !ok {
			fmt.Printf("Error: unknown label used: %s\n", instruct.args[0])
			return
		}

		w.SetWord(instruct.addr, opcode|labelAddr)
	}

	if err := w.Flush(); err != nil {
		fmt.Printf("Failed to write output: %s\n", err.Error())
	}
}

func strCmp(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if unicode.ToLower(rune(a[i])) != unicode.ToLower(rune(b[i])) {
			return false
		}
	}

	return true
}
