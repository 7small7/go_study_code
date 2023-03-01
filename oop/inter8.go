package main

import (
	"fmt"
	"reflect"
)

// 使用面向对象的知识，实现一个多台的关系
// 多态是指同一个方法，在不同的对象下，所产生的结果是不同的。
type oopPolymorphic interface {
	show1()
}

type Show1 struct {
}

func (s Show1) show1() {
	fmt.Println("show1 method")
}

type Show2 struct {
}

func (s Show2) show1() {
	fmt.Println("show2 method")
}

func main() {
	show1 := Show1{}
	show2 := Show2{}
	show3 := Show2{}
	show(show1)
	show(show2)
	if reflect.DeepEqual(show3, show2) {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
}

func show(polymorphic oopPolymorphic) {
	polymorphic.show1()
}
