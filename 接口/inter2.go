package main

import (
	"fmt"
)

// 接口断言
type interface2 interface {
	show()
}

type user1 struct {
	name string
	age  uint
}

func (u user1) show() {
	fmt.Println("姓名", u.name, "年龄", u.age)
}

func getShow(i interface{}) {
	switch t := i.(type) {
	case interface2:
		t.show() // 姓名 张三 年龄 12
	default:
		fmt.Println("无法实现断言判断", t)
	}
}

func main() {
	user := user1{name: "张三", age: 12}
	getShow(user)
}
