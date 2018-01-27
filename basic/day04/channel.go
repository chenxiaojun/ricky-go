package main

import (
	"fmt"
	//"time"
)

// 通道又叫channel，顾名思义，channel的作用就是在多线程之间传递数据的
//func main() {
//	// 创建无缓冲的channel
//	//chreadandwrite := make(chan int)
//	//chonlyread := make(<-chan int) // 创建只读的channel
//
//	//fatal error: all goroutines are asleep - deadlock! 这个错误的意思是线程陷入来死锁，程序无法向下执行
//	ch := make(chan int, 1)
//	ch <- 1
//	go func() { // 不会被执行，因为主线程优先执行，执行完就退出了
//		v := <-ch
//		fmt.Println(v)
//	}()
//	fmt.Println("2")
//}

func main() {
	ch := make(chan int, 1)
	go func() {
		v := <-ch
		fmt.Println("enter")
		fmt.Println(v)
	}()
	ch <- 3
	ch <- 8 // 如果把赋值 提到 子线程后面
	// 就无需让子程序休眠，原因是channel在主线程被赋值。 就会导致主线程阻塞。 直到channel中的子线程被取出
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	ch <- 9
	fmt.Println("2")
	//time.Sleep(1*time.Second)
}


