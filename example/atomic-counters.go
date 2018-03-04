package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"runtime"
)

func main() {
	// 定义一个64位无符号的整数，用来做计数器
	var ops uint64 = 0

	// 开通50个协程 并发执行任务
	for i := 0; i < 50; i++ {
		go func(){
			for {
				// 对计数器每隔1ms 进行一次加一操作
				atomic.AddUint64(&ops, 1)

				// 允许其它Go协程执行
				runtime.Gosched()
			}
		}()
	}

	// 让协程处理1s中
	time.Sleep(time.Second * 1)

	// 为了在计数器还在被其它Go协程更新时，安全的使用它，我们通过LoadUint64将当前值拷贝到opsFinal
	fmt.Println("ops: ", atomic.LoadUint64(&ops))
	fmt.Println("end")
}