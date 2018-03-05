package main

import "fmt"

// 通道是连接多个Go携程的通道
// 你可以从一个Go协程将值发送到通道，然后在别的Go协程中获取
func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	fmt.Println(<- messages)
}