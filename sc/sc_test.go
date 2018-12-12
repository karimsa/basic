// sc/sc_test.go
// Testing resets and counting.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package sc

import (
	"testing"
)

func TestCounter(t *testing.T) {
	for i := 0; i < 32; i++ {
		n := <-SC

		if n != i%16 {
			t.Fatalf("Expected %d, got %d\n", i%16, n)
		}
	}

	for i := 0; i < 5; i++ {
		n := <-SC
		if n != i {
			t.Fatalf("Expected %d, got %d\n", i, n)
		}
	}

	Reset()

	for i := 0; i < 5; i++ {
		n := <-SC
		if n != i {
			t.Fatalf("Expected %d, got %d\n", i, n)
		}
	}
}
