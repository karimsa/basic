// data/memory.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

import (
	"fmt"

	"github.com/karimsa/basic/constants"
)

type MemMode uint8

var (
	memory = make([]uint16, constants.MemorySize)
	DumpSize = 20
	memMode MemMode
)

const (
	// MemWrite writes to the bus
	MemWrite MemMode = iota
	
	// MemRead reads from the bus
	MemRead
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
func MemSelect(mode MemMode) {
	memMode = mode
}

func MemTick() {
	if memMode == MemRead {
		memory[AR.buffer] = BusRead()
	}

	memMode = MemWrite
}

func MemDump() {
	fmt.Printf("Memory[0:%d] => %#v\n", DumpSize, memory[0:DumpSize])
}
