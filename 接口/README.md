# 什么是接口

接口是一组方法的集合。接口只需要声明方法，而不需要声明方法体。其他的任意类型，都是可以实现接口中的方法。实现接口中的方法是隐式实现，
而不需要像其他面向对象语言，使用`implement`显示的指定实现关系。

# 声明方式

```go
package main

type demo interface {
	// 声明的方法，只要符合Go函数的定义即可。
	show()
	show1(age int) int
}
```

# 接口嵌套

1. 当前接口是允许直接嵌套其他的接口。

```go
package main

type interface7 interface {
	run()
}

type interface6 interface {
	show()
	interface7
}
```

> 要实现interface6，就必须实现show()和run()方法。

2. 如果当前的接口中声明的方法，在其他的接口中存在，默认会隐式实现其他的接口。

```go
package main

type demo2 interface {
	print()
}

type demo3 interface {
	print()
	show()
	kill()
}
```

> 上面的接口demo3种有方法print()，因此会被默认隐身嵌套了接口demo2。如果demo2中有其他的方法，但是在demo3中没有，则不会被隐式的嵌套。

# 断言

通过`v.(Type)`能够推断出当前类型所属什么类型，从而调用不同的实现方法。

```go
package main

import "fmt"

// 结构体签到
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
```