// cmd/boot/boot.go
// Boots up the computer with a target program.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package main

import (
	"os"
	"fmt"
	"time"

	"github.com/karimsa/basic/clock"
	"github.com/karimsa/basic/program"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: boot <program>\n")
		return
	}

	p, err := program.ReadProgram(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read program: %s\n", err.Error())
		return
	}

	p.Load()

	for {
		go clock.Tick()
		<-time.After(1*time.Second)
	}
}
