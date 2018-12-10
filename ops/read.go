// ops/read.go
// Reads an instruction and outputs an instruction type + address/opcode.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

const (
	MRI  uint16 = 0x0000
	IMRI uint16 = 0x0001
	RRI         = 0x7000
	IOI         = 0xF000
)

const (
	opcodeMask uint16 = 0xf000
)

func ReadInstruction(inst uint16) (uint16, uint16) {
	opType := inst & opcodeMask

	switch opType {
	case RRI:
		return opType, inst

	case IOI:
		return opType, inst

	default:
		// if opcode is 8-E, we should use
		// indirect addressing mode
		if opType > 0x6000 {
			return IMRI, inst &^ opcodeMask
		}

		// otherwise we use direct addressing
		return MRI, inst &^ opcodeMask
	}
}
