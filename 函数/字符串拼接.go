package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 1. 直接使用 "+"
	str1 := "你好"
	str2 := "世界"
	fmt.Println(str1 + str2) // 你好世界

	// 2. 使用strings.Join()
	slice1 := []string{str1, str2}
	fmt.Println(strings.Join(slice1, "")) // 你好世界

	// 3.使用fmt.Sprint()
	fmt.Println(fmt.Sprintf("%s%s", str1, str2)) // 你好世界

	// 4. 使用string.Builder()[推荐该方式]
	var str strings.Builder
	str.WriteString(str1)
	str.WriteString(str2)
	fmt.Println(str.String()) // 你好世界

	// 5. 通过bytes.Buffer()
	var byteStr bytes.Buffer
	byteStr.WriteString(str1)
	byteStr.WriteString(str2)
	fmt.Println(byteStr.String()) // 你好世界
}
