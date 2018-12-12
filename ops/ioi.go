// ops/ioi.go
// Input-output instructions.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

const (
	// INP inputs a character into AC
	INP uint16 = 0x0800

	// UT outputs a character from the AC
	UT = 0x0400

	// SKI skips next instruction on input flag
	SKI = 0x0200

	// SKO skips next instruction on output flag
	SKO = 0x0100

	// ION turns on the interrupt
	ION = 0x0080

	// IOP turns off the interupt
	IOP = 0x0040
)
