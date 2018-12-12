// data/registers.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

import (
	"fmt"
	"math"

	"github.com/karimsa/basic/alu"
	"github.com/karimsa/basic/debug"
	"github.com/karimsa/basic/ops"
)

type regMode uint8

const (
	none regMode = iota
	load
	inr
	clr
)

type Register struct {
	bitmask uint16
	buffer  uint16
	mode    regMode
}

func newRegister(size float64) *Register {
	return &Register{
		bitmask: uint16(math.Pow(2, size) - 1),
		buffer:  0,
		mode:    none,
	}
}

func (r *Register) Load() {
	r.mode = load
}

func (r *Register) Incr() {
	r.mode = inr
}

func (r *Register) Clr() {
	r.mode = clr
}

func (r *Register) set(v uint16) {
	r.buffer = v & r.bitmask
}

func (r *Register) tick(v uint16) {
	switch r.mode {
	case load:
		r.set(v)

	case inr:
		r.set(r.buffer + 1)

	case clr:
		r.buffer = 0
	}
}

var (
	// DR is the data register - it holds the memory operand
	DR = newRegister(16)

	// AC is the accumulator - the processor register
	AC = newRegister(16)

	// E is the single carry bit
	E = newRegister(1)

	// IR is the instruction register - it holds the instruction code
	IR = newRegister(16)

	// TR is the temporary register - for temp bits
	TR = newRegister(16)

	// AR is the address register - it holds the address for memory
	AR = newRegister(12)

	// PC is the program counter - it holds the address of the instruction
	PC = newRegister(12)

	// INPR is the input register - it holds the input char
	INPR = newRegister(8)

	// OUTR is the output register - it holds the output char
	OUTR = newRegister(8)
)

var (
	busRegisters = []*Register{
		// these registers read right off the bus
		DR, IR, TR, AR, PC, OUTR,

		// these are hard-wired to the ALU, so skipped
		// AC, E,

		// INPR is hard-wired to the input device
		// INPR,
	}
)

func ALUTick() {
	// though the clock is not actually connected to
	// the ALU, the shifting of the data in the registers
	// will cause the ALU content to change and therefore
	// it is a part of the 'RegTick' process
	ac, e := alu.Read(DR.buffer, uint8(INPR.buffer), AC.buffer)

	// manually tick on the AC - since it has selector
	// pins, we just tick and not set
	AC.tick(ac)
	AC.mode = none

	// set the value of E - it is not selected, always loads
	E.set(e)
}

func RegTick() {
	busValue := BusRead()

	for _, reg := range busRegisters {
		reg.tick(busValue)

		// OUTR is hard-wired to the output device, in this case, the
		// tty
		if reg == OUTR && reg.mode == load {
			if debug.Any {
				fmt.Printf("OUT: %s\n", string(byte(OUTR.buffer)))
			} else {
				fmt.Printf(string(byte(OUTR.buffer)))
			}
		}

		// auto-reset after pulse
		reg.mode = none
	}
}

// IR is readable by the control unit - since it is hard-wired
// to both the CU & the bus
func ReadIR() uint16 {
	return IR.buffer
}

// ShouldHalt signals the CU whether a halt instruction is
// loaded into the AR - mostly because I'm a bit fuzzy about the
// wiring for halting
func ShouldHalt() bool {
	return AR.buffer == ops.HLT
}

// TODO: Fuzzy about this wiring too - how does CU get the AR?
// there is no wiring there
func ReadAR() uint16 {
	return AR.buffer
}

func prettyMode(m regMode) string {
	switch m {
	case none:
		return "off"

	case load:
		return "load"

	case inr:
		return "incr"

	case clr:
		return "clr"

	default:
		return "unknown"
	}
}

func RegDump() {
	fmt.Printf("DR: %#v (%s)\n", DR.buffer, prettyMode(DR.mode))
	fmt.Printf("AC: %#v (%s)\n", AC.buffer, prettyMode(AC.mode))
	fmt.Printf("E: %#v (%s)\n", E.buffer, prettyMode(E.mode))
	fmt.Printf("IR: %#v (%s)\n", IR.buffer, prettyMode(IR.mode))
	fmt.Printf("TR: %#v (%s)\n", TR.buffer, prettyMode(TR.mode))
	fmt.Printf("AR: %#v (%s)\n", AR.buffer, prettyMode(AR.mode))
	fmt.Printf("PC: %#v (%s)\n", PC.buffer, prettyMode(PC.mode))
	fmt.Printf("INPR: %#v (%s)\n", INPR.buffer, prettyMode(INPR.mode))
	fmt.Printf("OUTR: %#v (%s)\n", OUTR.buffer, prettyMode(OUTR.mode))
}
