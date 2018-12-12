// control/control.go
// Combination circuit in CU.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package control

import (
	"fmt"
	"os"

	"github.com/karimsa/basic/data"
	"github.com/karimsa/basic/debug"
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
	op uint8
	I uint8
)

func Tick() {
	if debug.Control {
		fmt.Printf("SC => %d\n", sc.Read())
	}

	switch sc.Read() {
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
		op = uint8((data.ReadIR() & mid) >> 12)

		// AR <- IR(0-11)
		// no magic here to get rid of the MSBs - those
		// will be lost when converting from uint16 to uint12
		// in the register
		data.AR.Load()
		data.BusSelect(data.IR)

		// I <- IR(15)
		// shifted down to single bit
		I = uint8((data.ReadIR() & upper) >> 15)

	case 3:
		if op == 7 {
			if data.ShouldHalt() {
				fmt.Println("Halting")
				os.Exit(0)
			} else {
				// TODO: transfer the instruction to the ALU
				sc.Select(sc.CLR)
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
	}
}
