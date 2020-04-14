package main

import (
	"fmt"
	"sync"
)

type GroupWrapper struct {
	sync.WaitGroup
}

func (gw *GroupWrapper) Wrap(cb func()) {
	gw.Add(1)
	go func() {
		cb()
		gw.Done()
	}()
}

func test2() {
	var gw GroupWrapper
	for i := 0; i < 10; i++ {
		i := i
		gw.Wrap(func() {
			fmt.Println(i)
		})
	}
}
