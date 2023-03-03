package main

import "fmt"

//知识点 1： defer 的执行顺序。 采用栈的方式执行，先定义后执行。
//知识点 2：defer 与 return 谁先谁后。return 之后的语句先执行，defer 后的语句后执行
//知识点 3：函数的返回值初始化与 defer 间接影响。defer中修改了返回值得，实际返回的值是按照defer修改后的值进行返回。
//知识点 4：有名函数返回值遇见 defer 情况。
//知识点 5：defer 遇见 panic。按照defer的栈顺序，输出panic触发之前的defer。
//知识点 6：defer 中包含 panic。按照defer的栈顺序，输出panic触发之前的defer。并且defer中会接收到panic信息。
//知识点 7：defer 下的函数参数包含子函数。会先进行子函数的结果值，然后在按照栈的顺序进行输出。

func main() {
	fmt.Println(demo3())
}

func demo7() int {
	var a int
	defer func(a *int) {
		fmt.Println(*a)
		*a = 10
	}(&a)

	return 2
}

func function(index int, value int) int {
	fmt.Println(index)
	return index
}

func demo6() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
	// 3 4 2 1
}

func demo5() {
	defer func() {
		fmt.Println("1")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { fmt.Println("2") }()
	panic("panic")
	defer func() { fmt.Println("defer: panic 之后, 永远执行不到") }()
}

func demo4() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	panic("panic")
	defer func() {
		fmt.Println("3")
	}()
	defer func() {
		fmt.Println("4")
	}()
}

func demo3() (a int) {
	defer func() {
		fmt.Println(a)
		a = 3
	}()
	return 1
}

func demo2() int {
	defer func() {
		fmt.Println("2")
	}()
	return func() int {
		fmt.Println("1")
		return 4
	}()
}

func demo1() {
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()

	// 3 2 1
}
