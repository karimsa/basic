// main.go
// Entrypoint for basic computer.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package main

import (
	"fmt"

	"github.com/karimsa/basic/cpu"
	"github.com/karimsa/basic/program"
)

func main() {
	p, err := program.ReadProgram("./test/fixtures/hello-world.out")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	p.Load()
	cpu.Exec()
}
