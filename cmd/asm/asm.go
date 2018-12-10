// cmd/asm/asm.go
// Tiny assembler.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package main

import (
	"flag"
)

var (
	CommandLine = flag.NewFlagSet("asm", flag.ExitOnError)
	target      = CommandLine.String("out", "", "set the target file to write into")
)

func main() {
	if *target == "" {
		CommandLine.Usage()
		return
	}
}
