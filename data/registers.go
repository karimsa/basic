// cpu/registers.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

var (
	// DR is the data register - it holds the memory operand
	DR uint16

	// AC is the accumulator - the processor register
	AC uint16

	// E is the single carry bit - so using a bool to represent it
	E bool

	// IR is the instruction register - it holds the instruction code
	IR uint16

	// TR is the temporary register - for temp bits
	TR uint16

	// AR is the address register - it holds the address for memory
	// (actually 12-bits)
	AR uint16

	// PC is the program counter - it holds the address of the instruction
	// (actually 12-bits)
	PC uint16

	// INPR is the input register - it holds the input char
	INPR uint8

	// OUTR is the output register - it holds the output char
	OUTR uint8
)
