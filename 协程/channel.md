# 什么是channel

单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First
Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

# channel的操作

声明一个channel，需要使用`make()`为其分配内存，并初始化容量大小。如果没有使用`make()`分配内存，channel的值为nil。

```go
var channel名 chan channel类型
或者
channel名 := make(chan channel类型[, channel容量])
```

```go
package main

import "fmt"

func main() {
	var channel1 chan int
	fmt.Println(channel1) //<nil>

	channel := make(chan int, 1)
	defer close(channel)
	channel <- 10
	val, ok := <-channel
	fmt.Println(val, ok) // 10 true
}
```

声明了一个channel之后，一定要使用`close()`进行关闭channel。

## 写入数据

向channel中写入数据使用`channel <-`。

```go
channel <- "写入的值"
```

## 读取数据

向channel中读取数据使用`value, status : <- channel`。为channel中的值，ok表示当前channel是否处于打开状态，其值为`true`
或者`false`。

```go
val, ok := <-channel
```

# 缓冲区

## 无缓冲区

在了解channel阻塞情况，先看一下两个示例代码:

只有读端，没有写端。

```go
package main

func main() {
	channel := make(chan int)
	defer close(channel)
	<-channel
}
```

只有写端，没有读端。

```go
package main

func main() {
	channel := make(chan int)
	defer close(channel)
	channel <- 10
}
```

看到这里，先猜想一下两段代码的执行结果是怎么样的。通过执行，第一段代码的什么也不会输出，第二段代码将发生死锁。下面是第二段代码的输出错误信息。

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
/usr/local/var /www/go /go_study_notes/协程/channel1.go:7 +0x67
exit status 2
```

看到这里，你可能会纳闷为什么是这样的呢？这就需要了解到channel的缓冲区。这里总结一下关于缓冲区。

1. 无缓冲区的channel，默认是阻塞模式。只有读端和写端都准备好了，才不会发生阻塞。
2. 无缓冲区的channel，可以有只有读端，没有写端；但是不能只有写端，没有读端，否则会发生死锁。

## 有缓冲区的channel

有缓冲区，可以向channel中写入数据，直到容量被写满为止。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 5)
	defer close(channel)
	for i := 0; i < 5; i++ {
		channel <- i
	}
	for {
		val, ok := <-channel
		fmt.Println(val, ok)
		time.Sleep(time.Second * 4)
	}
}
```

上面的代码将输出如下的内容：

```go
0 true
1 true
2 true
3 true
4 true
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
/usr/local/var /www/go /go_study_notes/协程/channel1.go:15 +0xa8
exit status 2
```
