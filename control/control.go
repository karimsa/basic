// control/control.go
// Combination circuit in CU.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package control

import (
	"errors"
	"fmt"
	"os"

	"github.com/karimsa/basic/alu"
	"github.com/karimsa/basic/data"
	"github.com/karimsa/basic/debug"
	"github.com/karimsa/basic/ops"
	"github.com/karimsa/basic/sc"
)

const (
	// bit 15 of IR
	upper uint16 = 0x8000

	// bits 12-14 of IR
	mid uint16 = 0x7000

	// bits 0-11 of IR
	lower uint16 = 0x7ff0
)

// these are not registers, but they
// will not be used until the next clock
// pulse
var (
	op   uint8
	I    uint8
	halt bool
)

func Tick() {
	if debug.Control {
		fmt.Printf("SC => %d\n", sc.Read())
	}

	if halt {
		fmt.Println("Halting")
		os.Exit(0)
	}

	T := sc.Read()

	switch T {
	case 0:
		// AR <- PC
		data.AR.Load()
		data.BusSelect(data.PC)

	case 1:
		// IR <- M[AR]
		data.IR.Load()
		data.BusSelect(data.Memory)

		// PC <- PC + 1
		data.PC.Incr()

	case 2:
		// "decode" op from IR(12-14) - will just
		// and & shift it to get the 12-bits as decimal
		op = uint8((data.ReadRegister(data.IR) & mid) >> 12)

		// AR <- IR(0-11)
		// no magic here to get rid of the MSBs - those
		// will be lost when converting from uint16 to uint12
		// in the register
		data.AR.Load()
		data.BusSelect(data.IR)

		// I <- IR(15)
		// shifted down to single bit
		I = uint8((data.ReadRegister(data.IR) & upper) >> 15)

	case 3:
		if op == 7 {
			if I == 1 {
				if debug.Control {
					fmt.Println("IOI selected")
				}

				switch data.ReadRegister(data.AR) {
				case ops.INP:
					panic(errors.New("No input device is attached"))

				// OUTR <- AC
				case ops.UT:
					data.BusSelect(data.AC) // (bus) <- AC
					data.OUTR.Load()        // OUTR <- (bus)
					sc.Select(sc.CLR)       // SC <- 0

				// SKI skips next instruction on input flag
				case ops.SKI:
					panic(errors.New("instruction not implemented"))

				// SKO skips next instruction on output flag
				case ops.SKO:
					panic(errors.New("instruction not implemented"))

				// ION turns on the interrupt
				case ops.ION:
					panic(errors.New("instruction not implemented"))

				// IOP turns off the interupt
				case ops.IOP:
					panic(errors.New("instruction not implemented"))

				default:
					panic(fmt.Errorf("Unknown IOI operation at %#v: %#v", data.ReadRegister(data.PC)-1, data.ReadRegister(data.AR)))
				}
			} else {
				if debug.Control {
					fmt.Println("RRI selected")
				}

				switch data.ReadRegister(data.AR) {
				case ops.CLA:
					data.AC.Clr()     // AC <- 0
					sc.Select(sc.CLR) // SC <- 0

				case ops.CLE:
					data.E.Clr()      // E <- 0
					sc.Select(sc.CLR) // SC <- 0

				case ops.CMA:
					data.AC.Load()       // AC <- (alu)
					alu.Select(alu.COMP) // alu select COMP
					sc.Select(sc.CLR)    // SC <- 0

				case ops.CME:
					data.CompE()      // E <- E'
					sc.Select(sc.CLR) // SC <- 0

				case ops.CIR:
					data.AC.Load()       // AC <- (alu)
					alu.Select(alu.ASHR) // alu select SHR
					sc.Select(sc.CLR)    // SC <- 0

				case ops.CIL:
					data.AC.Load()       // AC <- (alu)
					alu.Select(alu.ASHL) // alu select SHL
					sc.Select(sc.CLR)    // SC <- 0

				case ops.INC:
					data.AC.Incr()    // AC <- AC + 1
					sc.Select(sc.CLR) // SC <- 0

				case ops.SPA:
					panic(errors.New("instruction not implemented"))

				case ops.SNA:
					panic(errors.New("instruction not implemented"))

				case ops.SZA:
					panic(errors.New("instruction not implemented"))

				case ops.SZE:
					panic(errors.New("instruction not implemented"))

				case ops.HLT:
					halt = true
					sc.Select(sc.CLR) // SC <- 0

				default:
					panic(fmt.Errorf("Unknown RRI instruction: %#v", data.ReadRegister(data.AR)))
				}
			}
		} else {
			if debug.Control {
				fmt.Printf("MRI selected\n")
			}

			// If I: AR <- M[AR]
			if I == 1 {
				data.AR.Load()
				data.BusSelect(data.Memory)
			}
		}

	case 4:
		// DR <- M[AR]
		if op != 7 {
			data.BusSelect(data.Memory)
			data.DR.Load()
		} else {
			panic(errors.New("Reached T4 for RRI/IOI execution cycles - should not happen"))
		}

	// MRI execution
	default:
		switch op {
		// AC <- DR & AC
		case ops.AND:
			data.AC.Load()      // AC <- (alu)
			alu.Select(alu.AND) // alu select AND
			sc.Select(sc.CLR)   // SC <- 0

		// AC <- DR + AC
		case ops.ADD:
			data.AC.Load()      // AC <- (alu)
			alu.Select(alu.AND) // alu select AND
			sc.Select(sc.CLR)   // SC <- 0

		// AC <- DR
		case ops.LDA:
			data.AC.Load() // AC <- (alu)
			alu.Select(alu.TR) // alu select TR
			sc.Select(sc.CLR)   // SC <- 0

		// M[AR] <- AC
		case ops.STA:
			data.MemSelect(data.MemRead)
			data.BusSelect(data.AC)

		// PC <- AR
		case ops.BUN:
			data.PC.Load()          // PC <- (bus)
			data.BusSelect(data.AR) // (bus) <- AR
			sc.Select(sc.CLR)       // SC <- 0

		// BSA:
		//   Save current address into M[AR]
		//   Branch to M[AR]+1
		case ops.BSA:
			if debug.Control {
				fmt.Printf("BSA\n")
			}

			switch T {
			// TR <- M[AR]
			case 5:
				data.TR.Load() // TR <- (bus)
				data.BusSelect(data.AR) // (bus) <- M[AR]

			// TR <- TR + 1, M[AR] <- PC
			case 6:
				data.TR.Incr() // TR <- TR + 1
				data.MemSelect(data.MemRead) // M[AR] <- (bus)
				data.BusSelect(data.PC) // (bus) <- PC

			// JMP to TR
			case 7:
				data.PC.Load()
				data.BusSelect(data.TR)
				sc.Select(sc.CLR)

			default:
				panic(fmt.Errorf("BSA reached T%d", T))
			}

		case ops.ISZ:
			switch T {
			// M[AR] <- M[AR] + 1
			case 5:
				data.BusSelect(data.Memory) // (bus) <- M[AR]
				data.TR.Load()              // TR <- (bus)

			case 6:
				data.TR.Incr() // TR <- TR + 1

			case 7:
				data.BusSelect(data.TR) // (bus) <- TR
				data.MemSelect(data.MemRead) // M[AR] <- (bus)

			case 8:
				// If M[AR] is zero: skip
				if data.ReadRegister(data.TR) == 0x0000 {
					data.PC.Incr()
				}
				sc.Select(sc.CLR)

			default:
				panic(fmt.Errorf("Reached T%d in ISZ execution", T))
			}

		default:
			panic(fmt.Errorf("Unknown MRI instruction: %#v", data.ReadRegister(data.AR)))
		}
	}
}
