# 切片定义

一片连续的内存空间加上长度与容量的标识。
其特点是容量和长度是动态增加，而不像数组一样是在声明时固定。
切片对元素的内容，与数组一致，都是在声明时指定，后期是不能进行修改的。

# 语法定义

```go
package main

import "fmt"

func main() {
	// 1. 通过数组的方式定义
	array := [3]int{1, 2, 3}
	slice := array[0:2]
	fmt.Println(slice)

	// 2. 直接声明定义
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice1)

	// 3. 通过make方式定义
	slice3 := make([]int, 2)
	slice3[0] = 1
	fmt.Println(slice3)

	// 4. 通过new方式定义[否]，这是因为new()返回的是某种数据类型的指针，虽然通过*指针能够获取到只指针的值，
	//  但通过new()初始化的切片是没有len和cap的，给切片进行复制时，就会出现panic。
	slice4 := new([]int)
	a := *slice4
	a[0] = 1
	fmt.Println(*slice4, a)
	// panic: runtime error: index out of range [0] with length 0
}
```

# 切片操作

```go
package main

import "fmt"

func main() {
	slice := make([]int, 5)
	// 1. 切片获取长度
	fmt.Println(len(slice))
	// 2. 切片获取容量
	fmt.Println(cap(slice)) // 默认情况下，在没有指定容量大小时，容量的大小=切片的长度

	// 3. 切片追加
	// Go默认提供append([]Type, ele ... Type)实现切片的追加，第一个参数必须是切片，第二个参数是具体的元素。
	// append()返回的是一个新的切片。
	// 3.1 默认是在尾部追加
	fmt.Println(append(slice, 1, 2))
	// 3.2 头部追加
	fmt.Println(append([]int{}, slice...))
	// 3.3 中间追加
	// 实现中间追加，可以使用append([]Type, ele ... Type)来实现
	fmt.Println(append(slice[0:3], 1))

	// 4. 获取指针信息
	fmt.Println(fmt.Sprintf("%p", slice))
}
```

# 切片扩容

切片的扩容演示的版本是`Go 1.18.6`。需要注意的是，不同的版本，在扩容的原理也是不同的。在 `Go 1.8`
之前，如果新切片是小于1024，新切片的容量直接是原切片的2倍，如果新切片的容量是大于1024，则按照1.25倍的大小进行扩容。

- 如果新切片的容量小于原切片的2倍，则扩大2倍；
- 如果新切片的容量小于256，则直接扩大2倍；
- 如果是大于256，则按照 新切片的容量 = 原切片的容量 + ( 原切片的容量 + 3 * 256) / 4 的规律进行动态扩容。

```go
package main

// Go扩容源码
func main() {
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		const threshold = 256
		if old.cap < threshold {
			newcap = doublecap
		} else {
			for 0 < newcap && newcap < cap {
				newcap += (newcap + 3*threshold) / 4
			}
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
}
```

演示代码：

```go
package main

import "fmt"

func main() {

	// 1. 新切片的容量小于老切片的2倍并且新切片的容量是小于256，则直接扩容2倍。
	slice1 := []int{1, 2}
	newSlice1 := append(slice1, 2)
	fmt.Println(len(slice1), len(newSlice1))  // 2, 3
	fmt.Println(cap(slice1), cap(newSlice1))  // 2, 4
	fmt.Println(fmt.Sprintf("%p", slice1))    // 0xc00001a0a0
	fmt.Println(fmt.Sprintf("%p", newSlice1)) // 0xc00001a0a0

	// 2. 如果新切片的容量大于原切片的两倍并且新切片的容量是小于256，则直接扩大2倍。
	slice2 := make([]int, 50)
	slice3 := make([]int, 60)
	newSlice2 := append(slice2, slice3...)
	fmt.Println(len(slice2), len(newSlice2))  // 50 110
	fmt.Println(cap(slice2), cap(newSlice2))  // 50 112
	fmt.Println(fmt.Sprintf("%p", slice2))    // 0xc000076000
	fmt.Println(fmt.Sprintf("%p", newSlice2)) // 0xc000078000

	// 3. 如果新切片的容量是大于256，则按照 新切片的容量 = 原切片的容量 + ( 原切片的容量 + 3 * 256) / 4 的规律进行动态扩容。
	slice4 := make([]int, 300)
	slice5 := make([]int, 100)
	newSlice3 := append(slice4, slice5...)
	fmt.Println(len(slice4), len(newSlice3))  // 300 400
	fmt.Println(cap(slice4), cap(newSlice3))  // 300 608
	fmt.Println(fmt.Sprintf("%p", slice4))    // 0xc0000bc000
	fmt.Println(fmt.Sprintf("%p", newSlice3)) // 0xc0000be000
}
```

> 在代码演示中，上述的规则并非是按照实际的情况得出结果，多少都有一点偏差。这是因为Go在处理切片扩容时，会考虑内存的对齐。关于内存对齐可以参考
> https://geektutu.com/post/hpg-struct-alignment.html。需要注意的是当切片进行扩容之后，可能会存在内存地址的变化，这说明新的切片在内存
> 中重新申请了一个内存空间，原有的内存空间将随着GC机制进行回收。

# 切片与数组的关系

```go
package main

import "fmt"

func main() {
	// 声明一个数组
	array1 := [4]int{1, 2, 3, 4}
	fmt.Println(fmt.Sprintf("%p", &array1)) // 0xc0000be000

	// 使用数组声明一个切片
	slice1 := array1[0:3]
	fmt.Println(fmt.Sprintf("%p", slice1)) // 0xc0000be000

	// 对切片进行扩容
	fmt.Println(fmt.Sprintf("%p", append(slice1, 1, 2, 3, 4, 5, 56, 6, 7, 8, 88, 8, 8, 8, 8, 8)))
	// 0xc00011e000
	// 对切片进行扩容之后，会发现内从地址也变了，这时因为原有的底层数组长度已无法满足新的切片长度。因此在内存中重新开辟了一段内存空间。
}
```

- 通过上述的代码演示，切片得到的内存地址和数组得到的内存地址是一致的。这也说明切片和数组执行的内存空间，都是同一片内存空间。
- 理论上来讲，切片其实底层也是数组的一种引用。换句话说，切片底层指向的就是一个数组。
- 通过切片的源码，其实不难发现，切片底层是一个`struct`结构体，内部的元素有`指向底层数组的指针`、`当前切片的长度`
  和`当前切片的容量`。
  ![](http://qiniucloud.qqdeveloper.com/202302191645409.png)

```go
package runtime

import (
	"unsafe"
)
# 切片定义

一片连续的内存空间加上长度与容量的标识。
其特点是容量和长度是动态增加，而不像数组一样是在声明时固定。
切片对元素的内容，与数组一致，都是在声明时指定，后期是不能进行修改的。

# 语法定义

```go
package main

import "fmt"

func main() {
	// 1. 通过数组的方式定义
	array := [3]int{1, 2, 3}
	slice := array[0:2]
	fmt.Println(slice)

	// 2. 直接声明定义
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice1)

	// 3. 通过make方式定义
	slice3 := make([]int, 2)
	slice3[0] = 1
	fmt.Println(slice3)

	// 4. 通过new方式定义[否]，这是因为new()返回的是某种数据类型的指针，虽然通过*指针能够获取到只指针的值，
	//  但通过new()初始化的切片是没有len和cap的，给切片进行复制时，就会出现panic。
	slice4 := new([]int)
	a := *slice4
	a[0] = 1
	fmt.Println(*slice4, a)
	// panic: runtime error: index out of range [0] with length 0
}
```

# 切片操作

```go
package main

import "fmt"

func main() {
	slice := make([]int, 5)
	// 1. 切片获取长度
	fmt.Println(len(slice))
	// 2. 切片获取容量
	fmt.Println(cap(slice)) // 默认情况下，在没有指定容量大小时，容量的大小=切片的长度

	// 3. 切片追加
	// Go默认提供append([]Type, ele ... Type)实现切片的追加，第一个参数必须是切片，第二个参数是具体的元素。
	// append()返回的是一个新的切片。
	// 3.1 默认是在尾部追加
	fmt.Println(append(slice, 1, 2))
	// 3.2 头部追加
	fmt.Println(append([]int{}, slice...))
	// 3.3 中间追加
	// 实现中间追加，可以使用append([]Type, ele ... Type)来实现
	fmt.Println(append(slice[0:3], 1))

	// 4. 获取指针信息
	fmt.Println(fmt.Sprintf("%p", slice))
}
```

# 切片扩容

切片的扩容演示的版本是`Go 1.18.6`。需要注意的是，不同的版本，在扩容的原理也是不同的。在 `Go 1.8`
之前，如果新切片是小于1024，新切片的容量直接是原切片的2倍，如果新切片的容量是大于1024，则按照1.25倍的大小进行扩容。

- 如果新切片的容量小于原切片的2倍，则扩大2倍；
- 如果新切片的容量小于256，则直接扩大2倍；
- 如果是大于256，则按照 新切片的容量 = 原切片的容量 + ( 原切片的容量 + 3 * 256) / 4 的规律进行动态扩容。

```go
package main

// Go扩容源码
func main() {
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		const threshold = 256
		if old.cap < threshold {
			newcap = doublecap
		} else {
			for 0 < newcap && newcap < cap {
				newcap += (newcap + 3*threshold) / 4
			}
			if newcap <= 0 {
				newcap = cap
			}
		}
	}
}
```

演示代码：

```go
package main

import "fmt"

func main() {

	// 1. 新切片的容量小于老切片的2倍并且新切片的容量是小于256，则直接扩容2倍。
	slice1 := []int{1, 2}
	newSlice1 := append(slice1, 2)
	fmt.Println(len(slice1), len(newSlice1))  // 2, 3
	fmt.Println(cap(slice1), cap(newSlice1))  // 2, 4
	fmt.Println(fmt.Sprintf("%p", slice1))    // 0xc00001a0a0
	fmt.Println(fmt.Sprintf("%p", newSlice1)) // 0xc00001a0a0

	// 2. 如果新切片的容量大于原切片的两倍并且新切片的容量是小于256，则直接扩大2倍。
	slice2 := make([]int, 50)
	slice3 := make([]int, 60)
	newSlice2 := append(slice2, slice3...)
	fmt.Println(len(slice2), len(newSlice2))  // 50 110
	fmt.Println(cap(slice2), cap(newSlice2))  // 50 112
	fmt.Println(fmt.Sprintf("%p", slice2))    // 0xc000076000
	fmt.Println(fmt.Sprintf("%p", newSlice2)) // 0xc000078000

	// 3. 如果新切片的容量是大于256，则按照 新切片的容量 = 原切片的容量 + ( 原切片的容量 + 3 * 256) / 4 的规律进行动态扩容。
	slice4 := make([]int, 300)
	slice5 := make([]int, 100)
	newSlice3 := append(slice4, slice5...)
	fmt.Println(len(slice4), len(newSlice3))  // 300 400
	fmt.Println(cap(slice4), cap(newSlice3))  // 300 608
	fmt.Println(fmt.Sprintf("%p", slice4))    // 0xc0000bc000
	fmt.Println(fmt.Sprintf("%p", newSlice3)) // 0xc0000be000
}
```

> 在代码演示中，上述的规则并非是按照实际的情况得出结果，多少都有一点偏差。这是因为Go在处理切片扩容时，会考虑内存的对齐。关于内存对齐可以参考
> https://geektutu.com/post/hpg-struct-alignment.html。需要注意的是当切片进行扩容之后，可能会存在内存地址的变化，这说明新的切片在内存
> 中重新申请了一个内存空间，原有的内存空间将随着GC机制进行回收。

# 切片与数组的关系

```go
package main

import "fmt"

func main() {
	// 声明一个数组
	array1 := [4]int{1, 2, 3, 4}
	fmt.Println(fmt.Sprintf("%p", &array1)) // 0xc0000be000

	// 使用数组声明一个切片
	slice1 := array1[0:3]
	fmt.Println(fmt.Sprintf("%p", slice1)) // 0xc0000be000

	// 对切片进行扩容
	fmt.Println(fmt.Sprintf("%p", append(slice1, 1, 2, 3, 4, 5, 56, 6, 7, 8, 88, 8, 8, 8, 8, 8)))
	// 0xc00011e000
	// 对切片进行扩容之后，会发现内从地址也变了，这时因为原有的底层数组长度已无法满足新的切片长度。因此在内存中重新开辟了一段内存空间。
}
```

- 通过上述的代码演示，切片得到的内存地址和数组得到的内存地址是一致的。这也说明切片和数组执行的内存空间，都是同一片内存空间。
- 理论上来讲，切片其实底层也是数组的一种引用。换句话说，切片底层指向的就是一个数组。
- 通过切片的源码，其实不难发现，切片底层是一个`struct`结构体，内部的元素有`指向底层数组的指针`、`当前切片的长度`
  和`当前切片的容量`。
  ![](http://qiniucloud.qqdeveloper.com/202302191645409.png)

```go
package runtime

import (
	"unsafe"
)

type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int            // 当前切片的长度
	cap   int            // 当前切片的容量
}
```

# 切片参数

- 切片属于引用类型，而不是值类型。
- 在Go语言中，不管是引用类型还是值类型，在函数传参都是属于值传递，只是底层指向的是同一个内存空间而已。
- 切片作为一个种引用类型，在传递参数时也属于值传递，只是新创建的变量的值是一个指针，该指针指向的是同一个底层数组。
  对该底层数组的改变，所有指向该底层数组的切片，其内部元素也会改变。
  ![](http://qiniucloud.qqdeveloper.com/202302191719190.png)
  ![](http://qiniucloud.qqdeveloper.com/202302191719229.png)
  ![](http://qiniucloud.qqdeveloper.com/202302191719281.png)

```go
package main

import "fmt"

func main() {
	array1 := [4]int{1, 2, 3, 4}
	slice1 := array1[0:3]

	show(slice1)
}

// 对切片的改变是否会影响原切片
func changeSlice(slice []int) {
	slice[0] = 8
}

func show(slice []int) {
	changeSlice(slice)
	sum := 0
	for _, value := range slice {
		sum += value
	}
	fmt.Println("sum = ", sum)
	// 如果没有调用changeSlice()函数，返回的结果是6。
	// 如果调用了changeSlice()函数，返回的结果是13。
}
```
type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int            // 当前切片的长度
	cap   int            // 当前切片的容量
}
```

# 切片参数

- 切片属于引用类型，而不是值类型。
- 在Go语言中，不管是引用类型还是值类型，在函数传参都是属于值传递，只是底层指向的是同一个内存空间而已。
- 切片作为一个种引用类型，在传递参数时也属于值传递，只是新创建的变量的值是一个指针，该指针指向的是同一个底层数组。
  对该底层数组的改变，所有指向该底层数组的切片，其内部元素也会改变。
  ![](http://qiniucloud.qqdeveloper.com/202302191719190.png)
  ![](http://qiniucloud.qqdeveloper.com/202302191719229.png)
  ![](http://qiniucloud.qqdeveloper.com/202302191719281.png)

```go
package main

import "fmt"

func main() {
	array1 := [4]int{1, 2, 3, 4}
	slice1 := array1[0:3]

	show(slice1)
}

// 对切片的改变是否会影响原切片
func changeSlice(slice []int) {
	slice[0] = 8
}

func show(slice []int) {
	changeSlice(slice)
	sum := 0
	for _, value := range slice {
		sum += value
	}
	fmt.Println("sum = ", sum)
	// 如果没有调用changeSlice()函数，返回的结果是6。
	// 如果调用了changeSlice()函数，返回的结果是13。
}
```