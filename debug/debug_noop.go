// debug/debug_noop.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

// +build !debug

package debug

const (
	Register = false
	Clock    = false
	CPU      = false
	ALU      = false
)
