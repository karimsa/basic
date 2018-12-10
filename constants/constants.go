// cpu/arch.go
// Architecture constants.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package constants

import (
	"math"
)

var (
	WordSize   = 16
	OpcodeSize = 4
	AddrSize   = WordSize - OpcodeSize

	MemorySize = int(math.Pow(2, float64(AddrSize)))
)
