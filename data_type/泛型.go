package main

import (
	"fmt"
	"time"
)

// SliceType 1. 定义泛型切片
type SliceType[T float32 | float64 | int] []T

// MapType 2. 定义泛型map
type MapType[key string, value any] map[key]value

// ChannelType 3. 泛型channel
type ChannelType[T any] chan T

func main() {
	slice := SliceType[int]{1, 3}
	fmt.Println(slice)
	fmt.Printf("%T", slice)

	mapType := MapType[string, string]{"key": "value"}
	mapType["key"] = "张三"
	fmt.Println(mapType)
	fmt.Printf("%T", mapType)
	fmt.Println()

	channelType := make(ChannelType[any], 1)
	go func() {
		if value, ok := <-channelType; ok {
			fmt.Println(value)
		}
	}()
	channelType <- 1
	time.Sleep(time.Second * 1)

	run("aaa")
	run(123)
	run(1.23)
	run(1.2)
}

// 定义一个泛型函数
func run[T string | float64 | float32 | int](a T) {
	fmt.Println(a)
}
