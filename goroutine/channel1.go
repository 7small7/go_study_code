package main

import "fmt"

// channel的声明
func main() {
	var channel1 chan int
	fmt.Println(channel1) //<nil>

	channel := make(chan int, 1)
	defer close(channel)
	channel <- 10
	val, ok := <-channel
	fmt.Println(val, ok) // 10 true
}
