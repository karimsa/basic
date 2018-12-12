// cpu/sc.go
// 4-bit sequence counter.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package sc

var (
	sc = make(chan int)

	min  = 0
	curr = 0
	max  = 15
)

// SC is a counter+decoder that will emit decimal values
// between 0 and 15
var SC <-chan int = sc

func Reset() {
	curr = -1
}

func init() {
	go func() {
		for {
			sc <- curr

			curr++
			if curr > max {
				curr = min
			}
		}
	}()
}
