package main

import "fmt"

// ping 函数定义了一个只允许发送数据的通道，尝试使用这个通道来接收数据将会得到一个编译时错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong函数允许通道pings来接收数据，另一通道pongs来发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
