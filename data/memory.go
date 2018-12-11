// data/memory.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

import (
	"github.com/karimsa/basic/constants"
)

var (
	memory = make([]uint16, constants.MemorySize)
)

type MemMode uint8

const (
	MemRead MemMode = iota
	MemWrite
)

// UnsafeMemWrite is only meant for program loading -
// it is to forcefully write to a piece of memory without going
// through the bus
func UnsafeMemWrite(pos int, word uint16) {
	memory[pos] = word
}

// Select changes the memory inputs to either load a word
// from the bus or switches the bus input to read off the
// memory (which simulates a write into the bus)
func Select(mode MemMode) {
	if mode == MemRead {
		BusSelect(7)
	} else {
		memory[AR.buffer] = BusRead()
	}
}

func Dump() []uint16 {
	return memory
}
