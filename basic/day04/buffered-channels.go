package main

import "fmt"

// channel可以是_带缓冲带。 为make提供第二个参数作为缓冲长度来初始化一个缓冲channel
// 向缓冲channel发送数据的时候，只有在缓冲区满的时候才会阻塞。 当缓冲区清空的时候接受阻塞。
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3  // all goroutines are asleep deadlock
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}
