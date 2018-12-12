// alu/adder.go
// 16-bit full adder.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package alu

import "fmt"

const (
	signMask int16 = -32768
)

func fullAdder(left int16, right int16) (int16, bool) {
	sum16 := left + right

	fmt.Printf("sign(%d) => %d\n", left, left&signMask)

	// if the operands have the same sign but the sum has a different one, there was an overflow
	if (left&signMask) == (right&signMask) && (left&signMask) != (sum16&signMask) {
		return sum16, true
	}

	// the second strategy is to check the bit addition on the last half-adder in the circuit,
	// but we don't have half-adders here - so instead will try shifting off the overflow bit
	// and then checking to see if it was lost

	sum32 := int32(left) + int32(right)
	return sum16, sum32 != (sum32 << 16 >> 16)
}
