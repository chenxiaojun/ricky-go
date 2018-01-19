package main

import (
	"fmt"
	"time"
)

// goroutine 是由Go运行环境管理的轻量级线程
// go f(x, y, z)
// 开启一个新的goroutine执行 f(x, y, z)
// f, x, y和z是当前goroutine中定义的，但是在新的goroutine中运行`f`
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("Hello")
}