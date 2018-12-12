// alu/alu.go
// Arithmetic logic unit.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package alu

import (
	"fmt"
)

type ALUOp uint8

var (
	mode      ALUOp
	usignMask uint16 = 0x8000
)

const (
	ADD ALUOp = iota
	SUB
	ASHL
	ASHR
	AND
	OR
	TR
	COMP
)

func Select(op ALUOp) {
	mode = op
}

func Read(DR uint16, INPR uint8, AC uint16) (uint16, uint16) {
	switch mode {
	case ADD:
		sum, overflow := fullAdder(int16(AC), int16(DR))
		if overflow {
			return uint16(sum), 1
		}
		return uint16(sum), 0

	case SUB:
		sum, overflow := fullAdder(int16(AC), -1*int16(DR))
		if overflow {
			return uint16(sum), 1
		}
		return uint16(sum), 0

	// arithmetic shift left - which keeps the sign bit
	case ASHL:
		sign := AC & usignMask
		shifted := AC << 1

		return shifted | sign, 0

	// arithmetic shift right - which keeps the sign bit
	case ASHR:
		sign := AC & usignMask
		shifted := AC >> 1

		return shifted | sign, 0

	case AND:
		return AC & DR, 0

	case OR:
		return AC | DR, 0

	case TR:
		return DR, 0

	case COMP:
		return ^AC, 0

	default:
		panic(fmt.Errorf("Unknown operation in ALU: %#v", mode))
	}
}
