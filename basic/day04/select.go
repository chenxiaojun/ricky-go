package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select 语句使得一个goroutine在多个通讯操作上等待
		// select会阻塞，知道条件分支中但某个可以继续执行，这是就会执行那个条件分支。 当多个都准备好的时候，会随机选择一个
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 创建10个线程 线程都是读出c里面的数据
		}
		quit <- 0 // push 0 into quit
	}()
	fibonacci(c, quit)
	fmt.Println("exit")
}

