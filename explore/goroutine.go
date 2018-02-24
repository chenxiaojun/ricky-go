package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
进程：进程包含了这个程序运行所需的各种资源，它像一个容器，是属于这个程序的工作空间。 比如它里面的内存空间，文件句柄，设备和线程等
线程：线程是一个执行的空间，线程会被操作系统调用，来在不同的处理器上运行编写的代码任务
进程包含多个线程，是一个比较大的概念。 如开启qq就是一个进程，qq传输文件，传输声音，视频又分别可划归为线程
一个进程在启动的时候，会有一个主线程，当这个主线程结束的时候，程序就终止了

Go中的并发指的是 让某个函数独立于其他函数运行的能力，一个goroutine就是一个独立的单元，goroutine运行时会在逻辑处理器上调度这些goroutine来运行
一个逻辑处理器绑定一个操作系统线程，所以说goroutine不是线程，它是一个协程。
*/

/**
并行：
	创建多个逻辑处理器，这样调度器就可以同时分配 全局运行队列 中的 goroutine 到不同的 逻辑处理上并行执行
*/
func main() {
	start := time.Now()
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2) // 设置一个计数器，这里计数器为2
	var count int

	go func() {
		defer wg.Done() // 当一个goroutine运行结束，就把计数器的值减1
		for i := 1; i < 1000000; i++ {
			count += i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 1000000; i++ {
			count += i
		}
	}()
	wg.Wait() // 当计数器的值被减少到0的时候，才会继续主线程。 如果是大于0，就会阻塞
	fmt.Println(count, time.Since(start).Seconds())
}
