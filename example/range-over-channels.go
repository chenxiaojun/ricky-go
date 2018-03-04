package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// 这里很关键
	close(queue)

	// 这个range迭代从queue中得到每个值。 因为我们在前面close了这个通道，所以这个
	// 迭代会在接收完2个值之后结束，如果我们没有close它，它将在这个循环中继续阻塞执行
	for elem := range queue {
		fmt.Println(elem)
	}

	test := make(chan string, 1) // 这里有1和没有1差别很大
	test <- "hi"
	fmt.Println(<-test)

	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2

	go func() {
		fmt.Println("coming")
		c2 <- 3 // 被阻塞
	}()

	time.Sleep(time.Second * 2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
