package main

import "fmt"

// 接口的基础定义
type interface1 interface {
	eat()
	run()
}

type People struct {
	name string
	age  uint
}

func (p *People) eat(name string) {
	p.name = name
}

func (p *People) run(age uint) {
	p.age = age
}

func main() {
	p := People{}
	p.eat("张三")
	p.run(12)
	fmt.Println(p.name, p.age) // 张三 12
}
