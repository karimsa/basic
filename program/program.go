// program/program.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package program

import (
	"fmt"
	"os"

	"github.com/karimsa/basic/constants"
	"github.com/karimsa/basic/data"
	"github.com/karimsa/basic/debug"
)

var (
	wordSize = constants.WordSize / 8
)

type Program struct {
	fd *os.File

	instructions [][2]uint16
	data         []uint16
}

type ProgramLine struct {
	label       *string
	instruction string
}

type ProgramWriter struct {
	fd *os.File

	labels  map[string]uint16
	lines   []ProgramLine
	maxaddr uint16
	data    []uint16
}

func ReadProgram(fpath string) (*Program, error) {
	fd, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	return &Program{
		fd: fd,
	}, nil
}

func WriteProgram(fpath string) (*ProgramWriter, error) {
	fd, err := os.OpenFile(fpath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}

	return &ProgramWriter{
		fd:     fd,
		labels: make(map[string]uint16),
		data:   make([]uint16, constants.MemorySize),
	}, nil
}

func (p *Program) Load() {
	addr := 0

	for {
		word, ok := p.readWord()
		if !ok {
			break
		}

		data.UnsafeMemWrite(addr, word)
		addr++
	}

	data.DumpSize = addr
	if debug.Any {
		fmt.Printf("Loaded program of %d words\n", data.DumpSize)
	}
}

func (p *Program) readWord() (uint16, bool) {
	word := make([]byte, wordSize)
	bytesRead, err := p.fd.Read(word)
	if bytesRead != wordSize || err != nil {
		return 0, false
	}

	return (uint16(word[0]) << 8) | uint16(word[1]), true
}

func (p *ProgramWriter) Flush() error {
	buff := make([]byte, wordSize)
	wordsWritten := 0

	// shrink data based on space used
	p.data = p.data[0 : p.maxaddr+1]
	if debug.ASM {
		fmt.Printf("data => %#v\n", p.data)
	}

	for _, word := range p.data {
		buff[0] = uint8(word >> 8) // upper 8-bits
		buff[1] = uint8(word)      // lower 8-bits

		bytesWritten, err := p.fd.Write(buff)
		if err != nil {
			return err
		}
		if bytesWritten < wordSize {
			return fmt.Errorf("Failed to write full word (%d) - only wrote %d bytes\n", wordSize, bytesWritten)
		}

		wordsWritten++
	}

	if debug.ASM {
		fmt.Printf("Stoping after writing %d words\n", wordsWritten)
	}
	return nil
}

func (p *ProgramWriter) SetWord(addr uint16, content uint16) {
	p.data[addr] = content
	if addr > p.maxaddr {
		p.maxaddr = addr
	}
}

func (p *ProgramWriter) SetLabel(label string, addr uint16) {
	p.labels[label] = addr
}

func (p *ProgramWriter) GetLabel(label string) (uint16, bool) {
	v, ok := p.labels[label]
	return v, ok
}

func (p *Program) close() {
	p.fd.Close()
}
