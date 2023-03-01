package main

import (
	"fmt"
	"strconv"
)

// 使用Go实现面向对象语言实现一个依赖注入

type Animal interface {
	eat()
	run()
	sleep()
}

type Dog struct {
	name  string
	age   int
	color string
}

type Cat struct {
	name  string
	age   int
	color string
}

func (d Dog) eat() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃东西。")
}

func (d Dog) run() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃运动。")
}

func (d Cat) sleep() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃谁家。")
}

func (d Cat) eat() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃东西。")
}

func (d Cat) run() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃运动。")
}

func (d Dog) sleep() {
	fmt.Println(d.name + "今年" + strconv.Itoa(d.age) + "岁了，有着" + d.color + "的毛发，平常爱吃谁家。")
}

func main() {
	dog := Dog{
		"小狗",
		12,
		"黄色",
	}
	cat := Cat{
		"小猫",
		12,
		"黄色",
	}
	// 实现依赖注入
	inject(dog)
	inject(cat)
}

func inject(a Animal) {
	a.eat()
	a.run()
	a.sleep()
}
