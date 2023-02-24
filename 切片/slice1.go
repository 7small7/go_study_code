package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	slice := make([]int, 5)
	// 1. 切片获取长度
	fmt.Println(len(slice))
	// 2. 切片获取容量
	fmt.Println(cap(slice)) // 默认情况下，在没有指定容量大小时，容量的大小=切片的长度

	// 3. 切片追加
	// Go默认提供append([]Type, ele ... Type)实现切片的追加，第一个参数必须是切片，第二个参数是具体的元素。
	// append()返回的是一个新的切片。
	// 3.1 默认是在尾部追加
	fmt.Println(append(slice, 1, 2))
	// 3.2 头部追加
	fmt.Println(append([]int{}, slice...))
	// 3.3 中间追加
	// 实现中间追加，可以使用append([]Type, ele ... Type)来实现
	fmt.Println(append(slice[0:3], 1))

	// 4. 获取指针信息
	fmt.Println(fmt.Sprintf("%p", slice))

	// 6. 切片排序
	slice2 := []int{2, 3, 1, 0, 8, 3}
	sort.Ints(slice2)
	fmt.Println(slice2)
}
