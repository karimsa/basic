// data/test_utils.go
// Utils for testing.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

// +build test

package data

func (r *Register) TestSet(value uint16) {
	r.set(value)
	ALUTick()
}

func (r *Register) TestRead() uint16 {
	return r.buffer
}
