### GO语言基础学习

今天是农历2018年的最后一天，在去年我开始学习go语言，毕业一年多
一直都是使用java开发有幸18年能接触到容器技术并接触到go语言。
为此专门想多练练语法和复习（学习）数据结构

#### 一些概念

*内置类型*：数值类型、字符串类型、布尔类型。这些类型本质上都是原始类型。对这些值进行增加删除时都会创建一个新值。

*引用类型*：切片、映射、通道、接口、函数类型。当申明上述类型的变量时，创建的变量被称作`标头`（heade）值。从技术细节上说字符串也是一种引用类型。每一个引用类型的标头值都包含了一个指向底层数据的结构的指针。每个引用类型还包含一组独特 的字段，用于管理底层数据结构。因为标头值是为复制而设计的，所以永远不需要共享一个引用 类型的值。

*方法集*：方法集定义了一组关联到给定类型的值或者指针的方法。定义方法时使用的接收者的类型决 定了这个方法是关联到值，还是关联到指针，还是两个都关联。

| values | methods receivers |
| ------ | ----------------- |
|   T    |       (t T)       |
|   *T   |	(t T) and (t *T)|

T 类型的值的方法集只包含值 接收者声明的方法。而指向 T 类型的指针的方法集既包含值接收者声明的方法，也包含指针接收 者声明的方法

`也就是指针类型的接受者实现了接口T类型是访问不到的`

*嵌入类型*：嵌入类型是将已有的类型直接声明在新的结构类型中，被嵌入的类型被称为新的外部类型的内部类。

*WatiGroup*: 是一个计数信号量，可以用来记录并维护运行的goroutine.如果WaitGroup的值大于0，Wait方法就会阻塞。
```go
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {	
		defer wg.Done()
	}
	
	wg.Wait()
}
```

*竞争状态*：如果两个或多个goroutine在没有相互同步的情况下，访问某个共享资源，并试图同时读和写这个资源，就处于相互竞争的状态

*无缓冲的通道*：是指在接受前没有能力保存任何值的通道

*有缓冲的通道*：不会要求goroutine之间必须同时完成发送和接受。只有在通道在没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。

#### golang中os/signal包的使用

golang中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。

监听全部信号：
```go
package main

import (
    "fmt"
    "os"
    "os/signal"
)

// 监听全部信号
func main()  {
    //合建chan
    c := make(chan os.Signal)
    //监听所有信号
    signal.Notify(c)
    //阻塞直到有信号传入
    fmt.Println("启动")
    s := <-c
    fmt.Println("退出信号", s)
}
```
```text
启动
go run example-1.go
启动

ctrl+c退出,输出
退出信号 interrupt

kill pid 输出
退出信号 terminated
```

监听指定信号
```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

// 监听指定信号
func main()  {
    //合建chan
    c := make(chan os.Signal)
    //监听指定信号 ctrl+c kill
    signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
    //阻塞直到有信号传入
    fmt.Println("启动")
    //阻塞直至有信号传入
    s := <-c
    fmt.Println("退出信号", s)
}
```

```text

启动
go run example-2.go
启动

ctrl+c退出,输出
退出信号 interrupt

kill pid 输出
退出信号 terminated

kill -USR1 pid 输出
退出信号 user defined signal 1

kill -USR2 pid 输出
退出信号 user defined signal 2
```
