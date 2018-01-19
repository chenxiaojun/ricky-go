package main

import (
	"fmt"
	//"time"
)

//func main() {
	// channel两种创建方式
	// unBufferChan := make(chan int)  // 无缓冲的channel, 如果使用channel之前没有make，会出现dead lock错误
	//var x chan int
	//go func() {
	//	x <- 1
	//}()
	//<-x

	// 无缓冲：发送和接收动作是同时发生。 如果没有goroutine读取channel(<-channel) 则发送者(channel<-)会一直阻塞
	// 缓冲：缓冲channel类似一个有容量的队列。当队列满的时候发送者会阻塞；当队列空的时候接收者会阻塞
//}

//func main() {
//	var c1 chan string = make(chan string)
//	go func() {
//		time.Sleep(time.Second * 2)
//		c1 <- "result 1"
//	}()
//	fmt.Println("c1 is", <-c1)
//}

//func main() {
//	fmt.Println("Begin doing something")
//	c := make(chan bool)
//	go func() {
//		fmt.Println("Doing something...")
//		close(c) // 促发这一事件
//	}()
//	<-c // 等待sub goroutine 完成事件
//	fmt.Println("Done!")
//}


//func main() {
//	fmt.Println("Begin doing something")
//	c := make(chan bool)
//	// go routine在channel c上没有数据可读，阻止线程进一步读取
//	go func() {
//		fmt.Println("Doing something...")
//	}()
//	<-c // 等待sub goroutine 完成事件
//	fmt.Println("Done!")
//}

//func main() {
//	fmt.Println("Begin doing something")
//	c := make(chan bool)
//	// go routine在channel c上没有数据可读，阻止线程进一步读取
//	go func() {
//		fmt.Println("Doing something...")
//		c<- true
//	}()
//	<-c // 等待sub goroutine 完成事件
//	fmt.Println("Done!")
//}

func worker(start chan bool, index int) {
	<- start
	fmt.Println("This is Worker:", index)
}

func main() {
	start := make(chan bool)
	for i := 1; i <= 5; i++ {
		go worker(start, i)
	}
	close(start)
	select{}
}









