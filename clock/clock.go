// clock/clock.go
// Simple clock.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package clock

import (
	"fmt"

	"github.com/karimsa/basic/control"
	"github.com/karimsa/basic/data"
	"github.com/karimsa/basic/debug"
	"github.com/karimsa/basic/sc"
)

func Tick() {
	if debug.Clock {
		fmt.Println("-------------------")
		fmt.Printf("before clock tick:\n")
		data.RegDump()
		data.BusDump()
		data.MemDump()
		fmt.Println("-------------------")
	}

	data.RegTick()
	data.ALUTick()
	control.Tick()
	sc.Tick()

	if debug.Clock {
		fmt.Println("-------------------")
		fmt.Printf("after clock tick:\n")
		data.RegDump()
		data.BusDump()
		data.MemDump()
		fmt.Println("-------------------")
	}
}
