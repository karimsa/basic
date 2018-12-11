// data/notest_utils.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

// +build !test

package data

import (
	"errors"
)

func (r *Register) TestSet(value uint16) {
	panic(errors.New("Test functions should not be used outside of tests"))
}

func (r *Register) TestRead() uint16 {
	panic(errors.New("Test functions should not be used outside of tests"))
}
