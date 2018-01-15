package main

import "fmt"

// Go 只有一种循环结构，也就是`for`循环
// 基本的for循环除了没有`()`之外(甚至强制不能使用它们)，看起来跟别的语言基本就是一样的了 而{}是必须的
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
