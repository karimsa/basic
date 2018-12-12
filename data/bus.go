// data/bus.go
// Simple 16-bit common bus.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

import (
	"fmt"
)

var (
	bus    uint16
	sel *Register
	Memory *Register
)

func BusRead() uint16 {
	return bus
}

func BusDump() {
	fmt.Printf("Bus: %#v (select: %#v)\n", bus, sel)
}

// BusSelect changes the mux inputs for the mux
// and therefore reads the content of a different register
// into the bus - since the bus is not connected to the clock,
// it does not tick and instead simply emit the content
// instantly
func BusSelect(reg *Register) {
	sel = reg

	if reg == Memory {
		bus = memory[AR.buffer]
	} else {
		bus = reg.buffer
	}
}
