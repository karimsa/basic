// cmd/boot/boot.go
// Boots up the computer with a target program.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package main

import (
	"fmt"
	"os"

	"github.com/karimsa/basic/clock"
	"github.com/karimsa/basic/program"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: boot <program>")
		return
	}

	p, err := program.ReadProgram(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Printf("Failed to read program: %s\n", err.Error())
		return
	}

	p.Load()

	for {
		clock.Tick()
	}
}
