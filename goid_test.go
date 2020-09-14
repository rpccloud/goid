// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package goid

import (
	"sync"
	"testing"
)

func TestGetGidPos(t *testing.T) {
	if getGidPos() < 0 {
		t.Fatalf("getGidPos error")
	}
	// make fake error
	temp := goroutinePrefix
	defer func() {
		goroutinePrefix = temp
	}()
	goroutinePrefix = "fake "
	if getGidPos() != -1 {
		t.Fatalf("getGidPos error")
	}
}

func TestGidFast(t *testing.T) {
	idMap := make(map[int64]bool)
	waitCH := make(chan bool)
	testCount := 1000
	mu := &sync.Mutex{}
	for i := 0; i < testCount; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			id := GoRoutineId()
			idMap[id] = true
			if id > 0 {
				waitCH <- true
			} else {
				waitCH <- false
			}
		}()
	}
	for i := 0; i < testCount; i++ {
		if !<-waitCH {
			t.Fatalf("GoRoutineId test error")
		}
	}
	if len(idMap) != testCount {
		t.Fatalf("GoRoutineId test error")
	}
}

func TestGidSlow(t *testing.T) {
	temp := gidPos
	defer func() {
		gidPos = temp
	}()
	gidPos = -1

	idMap := make(map[int64]bool)
	waitCH := make(chan bool)
	testCount := 1000
	mu := &sync.Mutex{}
	for i := 0; i < testCount; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			id := GoRoutineId()
			idMap[id] = true
			if id > 0 {
				waitCH <- true
			} else {
				waitCH <- false
			}
		}()
	}
	for i := 0; i < testCount; i++ {
		if !<-waitCH {
			t.Fatalf("GoRoutineId test error")
		}
	}
	if len(idMap) != testCount {
		t.Fatalf("GoRoutineId test error")
	}
}

func BenchmarkG(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GoRoutineId()
	}
}
