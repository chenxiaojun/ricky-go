package main

import (
	"fmt"
)

//func main() {
//	ch := make(chan int)
//	close(ch)
//	// 重复关闭会panic
//	// close(ch)
//
//	// 向关闭的channel发送数据会panic
//	//ch<- 8
//
//	// 向关闭的channel读取数据不会panic，读出但是其默认值
//	// fmt.Println(<-ch)
//
//	// 检查channel是否关闭了
//	_, ok := <-ch
//	fmt.Println(ok)
//}

//func main() {
//	x := make(chan int)
//	go func() {
//		x <- 2
//	}()
//	fmt.Print(<-x)
//}

func main() {
	c := make(chan int, 2)

	go func() {
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
	}()

	c <- 3
	fmt.Println("a")
	c <- 9
	fmt.Println("b")
	c <- 19
	fmt.Println("c")
}