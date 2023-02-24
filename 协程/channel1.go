package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 5)
	defer close(channel)
	for i := 0; i < 5; i++ {
		channel <- i
	}
	for {
		val, ok := <-channel
		fmt.Println(val, ok)
		time.Sleep(time.Second * 4)
	}
}
