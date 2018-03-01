package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // 能够读取到函数到局部变量
	fmt.Println(nextInt())
	fmt.Println(nextInt()) // 3 测试

	nextInt2 := intSeq()
	fmt.Println(nextInt2()) // 1 测试
}

func intSeq() func() (int, string) {
	i := 0
	s := "测试"
	return func() (int, string) {
		i++
		return i, s
	}
}