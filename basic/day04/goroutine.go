package main

import (
	"fmt"
	"time"
)

// 在golang里面，使用go这个关键字，后面再跟一个函数就可以创建一个线程
// 后面这个函数可以是已经写好的函数，也可以是一个匿名函数
//func main() {
//	go fmt.Println("1")
//	fmt.Println("2")
//}


/**
   这个函数只会打印2，在于go程序会优先执行主线程，主线程执行完以后，程序会立即退出，没有多余的时间去执行子线程
如果在程序的最后让主线程休眠1秒钟，那么程序就有足够的时间来执行子线程
*/
func main() {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
	// time.Sleep(1 * time.Second) // 主动让主程序休眠1秒，这样就有时间执行其它线程来
}