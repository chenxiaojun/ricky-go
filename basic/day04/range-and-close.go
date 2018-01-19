package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i :=0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	// 发送者可以close一个channel来表示再没有值会被发送了
	// 接收者可以通过赋值语句的第二个参数来测试channel是否被关闭
	// 当没有值可以接收并且channel已经被关闭 v, ok := <-ch   ok =false
	// 通常情况下，无需关闭channel。 只有在需要告诉接收者没有更多的数据的时候才有必要进行关闭，例如中断一个range
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
// n = 10 c = 0 x = 1 y = 1
// c = 1 x = 1 y = 2
// c = 1 x = 2 y = 3
// c = 2 x = 3 y = 5
