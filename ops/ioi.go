// ops/ioi.go
// Input-output instructions.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

const (
	// INP inputs a character into AC
	INP uint16 = 0xF800

	// UT outputs a character from the AC
	UT = 0xF400

	// SKI skips next instruction on input flag
	SKI = 0xF200

	// SKO skips next instruction on output flag
	SKO = 0xF100

	// ION turns on the interrupt
	ION = 0xF080

	// IOP turns off the interupt
	IOP = 0xF040
)
