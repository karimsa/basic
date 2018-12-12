// cpu/sc.go
// 4-bit sequence counter.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package sc

var (
	min  = 0
	curr = 0
	max  = 15

	mode SCMode = INR
)

type SCMode int

const (
	INR SCMode = iota
	CLR
)

func Select(m SCMode) {
	mode = m
}

func GetSC() int {
	return curr
}

func Tick() {
	switch mode {
	case INR:
		curr++
		if curr > max {
			curr = min
		}

	case CLR:
		curr = 0
	}

	mode = INR
}
