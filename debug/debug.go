// debug/debug.go
// Debug constants.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

// +build debug

package debug

// change these for debugging
const (
	Register = true
	Clock    = true
	Control  = true
	ALU      = true
	ASM      = true
	SC       = false
)

// do not change this
const Any = true
