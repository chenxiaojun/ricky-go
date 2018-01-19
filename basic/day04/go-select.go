package main

import (
	"fmt"
	"time"
)

// select的代码形式和switch非常相似, 不过select的case里的操作语句只能是[IO操作]
// 此示例里面的select会一直等待等到某个case语句完成，也就是等到成功从ch1或者ch2中读到数据，才会结束
func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int)
	go func() {
		time.Sleep(3*1e9) // sleep one second
		ch <- 1
		timeout <- true
	}()
	select {
	case <- ch:
		fmt.Println("in ch case")
	case <- timeout:
		fmt.Println("in timeout case")
	}
}

