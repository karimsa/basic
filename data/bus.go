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

func BusSelect(loc uint8) {
	switch loc {
	case 1:
		bus = AR

	case 2:
		bus = PC

	case 3:
		bus = DR

	case 4:
		bus = AC

	case 5:
		bus = IR

	case 6:
		bus = TR

	case 7:
		bus = memory[AR]
	}
}
