package main

import (
	"fmt"
	"sync"
)

var count = 0
var wg sync.WaitGroup

func main() {
	wg.Add(3)

	//for i := 0; i < 5; i++ {
	//	go read(i)
	//}

	for i := 0; i < 3; i++ {
		go write(i)
	}

	wg.Wait()
}

//func read(n int) {
//	defer wg.Done()
//	fmt.Printf("读取goroutine %d结束，值为: %d\n", n, count)
//}

func write(n int) {
	defer wg.Done()
	count++
	fmt.Printf("写goroutine %d结束，新值为：%d\n", n, count)
}
