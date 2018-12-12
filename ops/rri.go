// ops/rri.go
// Register Reference Instructions.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

const (
	// CLA clears the AC
	CLA uint16 = 0x0800

	// CLE clears the content of E
	CLE = 0x0400

	// CMA complements AC
	CMA = 0x0200

	// CME complements E
	CME = 0x0100

	// CIR circulates right AC & E
	CIR = 0x0080

	// CIL circulate left AC & E
	CIL = 0x0040

	// INC increments the AC
	INC = 0x0020

	// SPA skips the next instruction if AC is positive
	SPA = 0x0010

	// SNA skips the next instruction if AC is negative
	SNA = 0x0008

	// SZA skips the next instruction if AC is zero
	SZA = 0x0004

	// SZE skips the next instruction if E is zero
	SZE = 0x0002

	// HLT halts the computer
	HLT = 0x0001
)
