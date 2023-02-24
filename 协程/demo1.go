package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	var mux sync.RWMutex
	go func() {
		time.Sleep(time.Second * 4)
		count += 1
	}()
	if count > 0 {
		mux.Unlock()
	}
	fmt.Println(count)
}
