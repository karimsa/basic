// cpu/sc.go
// 4-bit sequence counter.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package sc

import (
	"fmt"

	"github.com/karimsa/basic/debug"
)

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

func Read() int {
	return curr
}

func Tick() {
	switch mode {
	case INR:
		curr++
		if debug.SC {
			fmt.Printf("Incrementing SC to %d\n", curr)
		}

		if curr > max {
			if debug.SC {
				fmt.Printf("SC exceeded max value, resetting to %d\n", min)
			}

			curr = min
		}

	case CLR:
		if debug.SC {
			fmt.Printf("Resetting SC\n")
		}
		curr = 0

	default:
		panic(fmt.Errorf("unknown SC mode: %d", mode))
	}

	mode = INR
}
