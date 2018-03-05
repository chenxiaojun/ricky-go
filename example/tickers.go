package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		// 无限循环去执行呢打点任务，直到ticker被关掉
		for t := range ticker.C {
			fmt.Println("ticket at", t)
		}
	}()

	// 休眠1600毫秒
	time.Sleep(time.Millisecond * 1600)
	// 停掉计时器
	ticker.Stop()
	fmt.Println("ticket stopped")
}
