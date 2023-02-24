package main

import "fmt"

// 使用断言模拟多驱动存储

type Redis interface {
	cache()
}

type Mysql interface {
	cache()
	mysql()
}

type MongoDb interface {
	cache()
	mongodb()
}

type cache struct {
	host string
	port uint32
}

func (c cache) cache() {
	fmt.Println("我是一个cache方法")
}

//func (c cache) mysql() {
//	fmt.Println("我是一个mysql方法")
//}

func cacheType(i interface{}) {
	switch i.(type) {
	case MongoDb:
		fmt.Println("实现的接口是MongoDB")
	case Mysql:
		fmt.Println("实现的接口是MySQL")
	case Redis:
		fmt.Println("实现的接口是Redis")
	default:
		fmt.Println("还未实现任何接口")
	}
}

func main() {
	c := cache{host: "127.0.0.1", port: 6479}
	cacheType(c) //实现的接口是Redis
}
