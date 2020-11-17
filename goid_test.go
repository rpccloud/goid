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
			id := GetRoutineId()
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
			t.Fatalf("GetRoutineId test error")
		}
	}
	if len(idMap) != testCount {
		t.Fatalf("GetRoutineId test error")
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
			id := GetRoutineId()
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
			t.Fatalf("GetRoutineId test error")
		}
	}
	if len(idMap) != testCount {
		t.Fatalf("GetRoutineId test error")
	}
}

func BenchmarkGetRoutineId(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GetRoutineId()
		}
	})
}
