// utils/waitgroup.go
// WaitGroup wrapper.
//
// Copyright (C) 2018-present Karim Alibhai. All rights reserved

package utils

import (
	"sync"
)

type WaitGroup struct {
	wg *sync.WaitGroup
}

func (w *WaitGroup) Add(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		fn()
	}()
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}
