// data/bus.go
// Simple 16-bit common bus.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package data

var (
	bus uint16
)

func BusRead() uint16 {
	return bus
}

// BusSelect changes the mux inputs for the mux
// and therefore reads the content of a different register
// into the bus - since the bus is not connected to the clock,
// it does not tick and instead simply emit the content
// instantly
func BusSelect(loc uint8) {
	switch loc {
	case 1:
		bus = AR.buffer

	case 2:
		bus = PC.buffer

	case 3:
		bus = DR.buffer

	case 4:
		bus = AC.buffer

	case 5:
		bus = IR.buffer

	case 6:
		bus = TR.buffer

	case 7:
		bus = memory[AR.buffer]
	}
}
