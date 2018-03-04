package main

import (
	"fmt"
	"time"
)

// 速率控制
func main() {
	jobs := make(chan int, 10)
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//tick := time.Tick(time.Millisecond * 500)
	ticker := time.NewTicker(time.Millisecond * 500)
	for job := range jobs {
		//<-tick
		<-ticker.C
		fmt.Printf("job: %d, current: %v\n", job, time.Now())
	}

	limiter := make(chan time.Time, 3)
	for i := 0; i < 2; i++ {
		limiter <- time.Now()
	}
	//close(limiter)

	// 因为limiter的通道没有关闭掉，所有当上面limiter值取完，就会来执行这个协程里面的
	// 这里面每个1000毫米给limiter塞进一个值，从而达到下面的阻塞作用
	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			limiter <- t
		}
	}()

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}
