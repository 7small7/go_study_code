package main

import "fmt"

// 结构体嵌套
type interface3 interface {
	show()
}

type interface4 interface {
	show()
}

type str2 struct {
	name string
	age  uint
}

func (s str2) show() {
	fmt.Println("姓名", s.name, "年龄", s.age)
}

func getType(i interface{}) {
	switch i.(type) {
	case interface4:
		fmt.Println("interface4")
	case interface3:
		fmt.Println("interface3")
	default:
		fmt.Println("unknow type")
	}
}

func main() {
	str2 := str2{name: "张三", age: 12}
	getType(str2) // interface4, 结构体都实现了interface3和interface4接口，在switch判断，谁在前则优先执行谁。
}
