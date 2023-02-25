package main

import "fmt"

// bool类型
func main() {
	var b1 bool
	fmt.Println(b1) // false
	b1 = false
	fmt.Println(b1) // false

	b2 := true
	fmt.Println(b2) // true

	b3 := 1
	if b3 > 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
