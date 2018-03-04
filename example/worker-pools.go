package main

import (
	"fmt"
	"time"
)

// 线程池
func main() {
	start := time.Now()
	// 用来发送数据
	jobs := make(chan int, 100)
	// 用来接收结果
	results := make(chan int, 100)

	//启动3个works来处理jobs
	for i := 1; i < 4; i++ {
		go works(i, jobs, results)
	}

	// 给jobs发送任务
	for i := 0; i <= 9; i++ {
		jobs <- i
	}
	// jobs的任务已经发送完毕，关闭jobs
	close(jobs)

	for i := 0; i <= 9; i++ {
		fmt.Println(<-results)
	}
	fmt.Printf("Process cost %.2fs\n", time.Since(start).Seconds())
}

func works(work int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("workid: %d, jvalue: %d\n", work, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}
