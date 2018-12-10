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

func MemWrite() {
	memory[AR] = BusRead()
}

func MemRead() {
	BusSelect(7)
}

func Dump() []uint16 {
	return memory
}
