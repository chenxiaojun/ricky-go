package main

import (
	"sync"
	"fmt"
	"time"
)

/**
	控制并发有两种方式，一种是WaitGroup 另外一种是Context
	WaitGroup 它是一种控制并发的方式，控制多个goroutine同时完成
 */
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()

	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("END")
}