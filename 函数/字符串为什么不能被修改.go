package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 常见的字符串重新赋值(这其实不是修改，而是一种重新赋值)
	str1 := "你好"
	str1 = "我不好"
	fmt.Println(str1, str1[0], len(str1), utf8.RuneCountInString(str1)) // 我不好 230 9 3

	// 循环字符串
	for _, i2 := range str1 {
		fmt.Println(i2) // 打印对应的字节码
	}
	//25105
	//19981
	//22909

	// 修改字符串某一个字符串，这种方式是不允许的。
	//str1[0] = "你好"
	// Go中的字符串的数据结构体是由一个指针和长度组成的结构体，该指针指向的一个切片才是真正的字符串值。底层其实是一个指针和当前字符串的长度组成。
	// 指针指向的是[]byte类型的数组。
	//type stringStruct struct {
	//	str unsafe.Pointer
	//	len int
	//}

	b := []byte(str1)
	fmt.Println(b) // 230 136 145 228 184 141 229 165 189

	byteSlice := []byte{230, 136, 145, 228, 184, 141, 229, 165, 189}
	fmt.Println(string(byteSlice)) // 我不好
	byteSlice[2] = 108
	fmt.Println(string(byteSlice)) // �l不好
}
