// ops/read_test.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package ops

import (
	"testing"
)

func TestMRIReads(t *testing.T) {
	if opt, opcode := ReadInstruction(0x23AF); opt != MRI || opcode != 0x3AF {
		t.Fatalf("Failed to break down 0x23AF into 0x2 & 0x03AF (got: opt=%#v, opcode=%#v)", opt, opcode)
	}

	if opt, opcode := ReadInstruction(0xB93A); opt != IMRI || opcode != 0x93A {
		t.Fatalf("Failed to break down 0xB93A into 0xB & 0x93A (got: opt=%#v, opcode=%#v)", opt, opcode)
	}
}

func TestRRIReads(t *testing.T) {
	if opt, opcode := ReadInstruction(0x7800); opt != RRI || opcode != CLA {
		t.Fatalf("Failed to break down 0x7800 into 0x7 & 0x0800 (got: opt=%#v, opcode=%#v)", opt, opcode)
	}

	if opt, opcode := ReadInstruction(0x7008); opt != RRI || opcode != SNA {
		t.Fatalf("Failed to break down 0x7008 into 0x7 & 0x0008 (got: opt=%#v, opcode=%#v)", opt, opcode)
	}
}

func TestIOIReads(t *testing.T) {
	if opt, opcode := ReadInstruction(0xF400); opt != IOI || opcode != UT {
		t.Fatalf("Failed to break down 0xF400 into 0xF & 0x0400 (got: opt=%#v, opcode=%#v)", opt, opcode)
	}

	if opt, opcode := ReadInstruction(0xF080); opt != IOI || opcode != ION {
		t.Fatalf("Failed to break down 0xF080 into 0xF & 0x0080 (got: opt=%#v, opcode=%#v)", opt, opcode)
	}
}
