// data/reg_test.go
// Testing manual register function.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package test

import (
	"testing"

	"github.com/karimsa/basic/clock"
	"github.com/karimsa/basic/data"
)

func TestTransfer(t *testing.T) {
	// DR <- AC

	// for testing
	data.AC.TestSet(0x35ac)

	// 1) Tell DR to load
	data.DR.Load()

	// 2) Update bus mux to read AC
	data.BusSelect(data.AC)

	// AC should have its own value
	if data.AC.TestRead() != 0x35ac {
		t.Fatalf("AC is holding: %#v (should be 0x35ac)", data.AC.TestRead())
	}

	// DR should have zero
	if data.DR.TestRead() != 0x0000 {
		t.Fatalf("DR is holding: %#v (should be zero)", data.DR.TestRead())
	}

	// 3) Emit pulse
	clock.Tick()

	// DR should now have the AC value in it
	if data.DR.TestRead() != 0x35ac {
		t.Fatalf("DR is holding: %#v, AC is holding: %#v (should be same)", data.DR.TestRead(), data.AC.TestRead())
	}
}

func TestSwap(t *testing.T) {
	// DR <- AC, AC <- DR

	initAC := uint16(0x35ac)
	initDC := uint16(0xf476)

	// for testing
	data.AC.TestSet(initAC)
	data.DR.TestSet(initDC)

	// 1) Tell both registers to load
	data.AC.Load()
	data.DR.Load()

	// 2) Update bus mux to read AC
	data.BusSelect(data.AC)

	// AC should have its own value
	if data.AC.TestRead() != initAC {
		data.RegDump()
		t.Fatalf("AC is holding: %#v (should be %#v)", data.AC.TestRead(), initAC)
	}

	// DR should have its initial value
	if data.DR.TestRead() != initDC {
		data.RegDump()
		t.Fatalf("DR is holding: %#v (should be %#v)", data.DR.TestRead(), initDC)
	}

	// 3) Emit pulse
	clock.Tick()

	// DR should have the AC value
	if data.DR.TestRead() != initAC {
		data.RegDump()
		t.Fatalf("DR <- AC; DR is holding: %#v (should be %#v)", data.DR.TestRead(), initAC)
	}

	// AC should have the DC value
	if data.AC.TestRead() != initDC {
		data.RegDump()
		t.Fatalf("AC <- DR failed; AC is holding: %#v (should be %#v)", data.AC.TestRead(), initDC)
	}
}
