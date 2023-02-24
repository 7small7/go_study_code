package main

import "fmt"

// 结构体嵌套
type interface5 interface {
	show()
}

type interface7 interface {
	run()
}

type interface6 interface {
	show()
	interface7
}

type str3 struct {
	name string
	age  uint
}

func (s str3) show() {
	fmt.Println("姓名", s.name, "年龄", s.age)
}

func GetType(i interface{}) {
	switch i.(type) {
	case interface5:
		fmt.Println("interface5")
	case interface6:
		fmt.Println("interface6")
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	str2 := str3{name: "张三", age: 12}
	GetType(str2) // interface5, 因为interface6继承了interface7，因此要实现interface6，则必须也实现interface7中的run()函数。
}
