// ops/rri.go
// Register Reference Instructions.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

const (
	// CLA clears the AC
	CLA uint16 = 0x7800

	// CLE clears the content of E
	CLE = 0x7400

	// CMA complements AC
	CMA = 0x7200

	// CME complements E
	CME = 0x7100

	// CIR circulates right AC & E
	CIR = 0x7080

	// CIL circulate left AC & E
	CIL = 0x7040

	// INC increments the AC
	INC = 0x7020

	// SPA skips the next instruction if AC is positive
	SPA = 0x7010

	// SNA skips the next instruction if AC is negative
	SNA = 0x7008

	// SZA skips the next instruction if AC is zero
	SZA = 0x7004

	// SZE skips the next instruction if E is zero
	SZE = 0x7002

	// HLT halts the computer
	HLT = 0x7001
)
