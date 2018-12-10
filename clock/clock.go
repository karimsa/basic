// clock/clock.go
// Simple clock.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package clock

import (
	"github.com/karimsa/basic/alu"
	"github.com/karimsa/basic/utils"
)

func Tick() {
	wg := utils.WaitGroup{}
	wg.Add(func() {
		alu.Tick()
	})
	wg.Wait()
}
