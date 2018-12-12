// program/program_test.go
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package program

import (
	"testing"
)

func test(fpath string, t *testing.T) {
	program, err := ReadProgram(fpath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if word, ok := program.readWord(); !ok {
		t.Fatalf("Failed to read first word from program")
	} else if word != 0x7800 {
		t.Fatalf("Read wrong first word, got: %#v (expected 0x7800)", word)
	}

	if word, ok := program.readWord(); !ok {
		t.Fatalf("Failed to read second word from program")
	} else if word != 0xc152 {
		t.Fatalf("Read wrong second word, got: %#v (expected 0xc152)", word)
	}

	if word, ok := program.readWord(); ok {
		t.Fatalf("Read non-existent third word: %#v", word)
	}
}

func TestProgramWriting(t *testing.T) {
	program, err := WriteProgram("../test/fixtures/test2.out")
	if err != nil {
		t.Fatalf(err.Error())
	}

	program.SetWord(0, 0x7800)
	program.SetWord(1, 0xc152)
	program.Flush()

	test("../test/fixtures/test2.out", t)
}

func TestProgramReading(t *testing.T) {
	test("../test/fixtures/test1.out", t)
}
