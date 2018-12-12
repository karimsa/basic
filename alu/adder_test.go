// alu/adder_test.go
// Tests for full adder.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package alu

import (
	"testing"
)

var (
	maxSafe16 = int16(32767)
	minSafe16 = int16(-32768)
)

func TestPosPlusPosOverflow(t *testing.T) {
	sum, overflow := fullAdder(maxSafe16, 1)
	if !overflow {
		t.Fatalf("sum => %#v, overflow => %#v\n", sum, overflow)
	}
}

func TestPosPlusPosNoOverflow(t *testing.T) {
	sum, overflow := fullAdder(maxSafe16-1, 1)
	if overflow || sum != maxSafe16 {
		t.Fatalf("sum => %#v, overflow => %#v\n", sum, overflow)
	}
}

func TestNegPlusPosNoOverflow(t *testing.T) {
	sum, overflow := fullAdder(minSafe16, 1)
	if overflow || sum != (minSafe16+1) {
		t.Fatalf("sum => %#v, overflow => %#v\n", sum, overflow)
	}
}

func TestNegPlusNegNoOverflow(t *testing.T) {
	sum, overflow := fullAdder(minSafe16+1, -1)
	if overflow || sum != minSafe16 {
		t.Fatalf("sum => %#v, overflow => %#v\n", sum, overflow)
	}
}

func TestNegPlusNegOverflow(t *testing.T) {
	sum, overflow := fullAdder(minSafe16, -1)
	if !overflow {
		t.Fatalf("sum => %#v, overflow => %#v\n", sum, overflow)
	}
}
