package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := []int{1, 2}
	slice2 := []int{1, 3}
	slice3 := []string{"1", "2"}
	//if slice1 == slice2 { // 直接会发生编译错误,invalid operation: slice1 == slice2 (slice can only be compared to nil)
	//	fmt.Println("==")
	//}

	if reflect.DeepEqual(slice1, slice2) { // != 数据类型相同，但值不相同
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}

	if reflect.DeepEqual(slice1, slice3) { // != ，数据类型不同，并且值也不同
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
}
