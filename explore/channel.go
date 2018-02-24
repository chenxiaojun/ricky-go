package main

import (
	"fmt"
	"os"
	"time"
)

//func main() {
//	// 通道使用chan关键字 并指定发送和接收的类型
//	// 通道和Map类型一样，可以使用内置的make函数声明初始化
//	ch := make(chan int)
//
//	go func() {
//		var sum int = 0
//		for i := 0; i < 10; i++ {
//			sum += i
//		}
//		ch <- sum
//	}()
//
//	fmt.Println(<-ch)
//	fmt.Println("End")
//}

//func main() {
//	one := make(chan int)
//	two := make(chan int)
//
//	go func() {
//		fmt.Println("coming")
//		one <- 100
//	}()
//
//	go func() {
//		fmt.Println("coming2")
//		v := <- one
//		two<-v
//	}()
//
//	fmt.Println(<-two) // coming2 coming 100
//}

//func main() {
//	ch := make(chan string)
//
//	// 假设有一批用户排队付款 要求按顺序给用户发送信息
//	users := []string {
//		"赵1", "钱2", "孙3", "李4", "周5", "吴6", "郑7", "王8",
//	}
//
//	for _, user := range users {
//		go func() {
//			ch <- user
//		}()
//	}
//
//	for _, msg := range <-ch {
//		fmt.Println(msg)
//	}
//}

//func main() {
//	// 创建容量为3，有缓冲的通道。
//	// 队列满的时候 发送操作会阻塞，  队列为空的时候，接受操作会阻塞
//	for i := 1; i < 10; i++ {
//		fmt.Println(test())
//	}
//}
//
//func test() string {
//	responses := make(chan string, 3)
//	go func() {responses <- a()}()
//	go func() {responses <- b()}()
//	go func() {responses <- c()}()
//	return <-responses
//}
//
//func a() string {
//	return "a"
//}
//
//func b() string {
//	return "b"
//}
//
//func c() string {
//	return "c"
//}

//func main() {
//	ch := make(chan string)
//
//	go func() {
//		ch <- "Ricky"
//		ch <- "test"
//	}()
//
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//	fmt.Println("haha")
//}

// 有时候，我们限制一个通道只可以接收，但是不能发送
// 有时候限制一个通道只可以发送，不可以接收   <-  单向通道

//func main() {
//	var send chan <- int // 只能发送
//	var receive <- chan int //只能接收
//}
