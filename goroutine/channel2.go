package main

import "time"

// 声明一个只读或者只写的channel
func main() {
	go onlyWrite()
	go onlyRead()
	time.Sleep(time.Second)
}

func onlyWrite() {
	ch1 := make(chan<- int, 1)
	defer close(ch1)

	//<-ch1
	// invalid operation: cannot receive from send-only channel ch1 (variable of type chan<- int)

}

func onlyRead() {
	//ch1 := make(<-chan int, 1)
	//ch1 <- 1
	// invalid operation: cannot send to receive-only channel ch1 (variable of type <-chan int)
}
