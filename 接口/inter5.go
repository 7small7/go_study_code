package main

type demo1 interface {
	print()
	play()
}

type demo2 interface {
	print()
}

type demo3 interface {
	print()
	show()
	kill()
}

// 1. 如果当前interface中的方法在其他的interface种定义，并且其他的中只有当前被隐式继承的
