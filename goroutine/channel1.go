package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		val, ok := <-channel
		fmt.Println(val, ok)
	}()

	go func() {
		channel <- 1
	}()

	defer close(channel)

	time.Sleep(time.Second)
}
