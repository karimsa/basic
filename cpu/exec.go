// cpu/exec.go
// Executes a program from memory.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package cpu

import (
	"fmt"

	"github.com/karimsa/basic/data"
)

func Exec() {
	fmt.Printf("mem => %+v\n", data.Dump())
}
