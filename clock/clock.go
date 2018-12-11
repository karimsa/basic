// clock/clock.go
// Simple clock.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package clock

import (
	"fmt"

	"github.com/karimsa/basic/data"
	"github.com/karimsa/basic/debug"
)

func Tick() {
	if debug.Clock {
		fmt.Println("-------------------")
		fmt.Printf("before clock tick:\n")
		data.RegDump()
		fmt.Println("-------------------")
	}

	data.RegTick()

	if debug.Clock {
		fmt.Println("-------------------")
		fmt.Printf("after clock tick:\n")
		data.RegDump()
		fmt.Println("-------------------")
	}
}
