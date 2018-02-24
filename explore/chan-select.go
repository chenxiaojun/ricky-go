package main

import (
	"fmt"
	"time"
)

/**
	一个goroutine启动，我们无法控制它，大部分是等待自我结束
	一直傻瓜式的办法是全局变量，其他地方通过修改这个变量完成结束通知，然后后台的goroutine不停的检查这个变量
如果发现被通知关闭了，就自我结束。 一般我们使用chan+select来实现这种
*/

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}
