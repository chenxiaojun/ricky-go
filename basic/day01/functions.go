package main

import (
	"fmt"
)

// 函数可以没有参数或接受多个参数
// add函数接受两个int类型的参数，返回值是int类型
func add(x int, y int) int {
	return x + y
}

// 当两个或多个连续的函数 函数命名的参数是同一类型，则除了最后面一个类型之外，其他都可以省略
func plus(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(3, 6))

	fmt.Println(plus(1, 67))

	fmt.Println(swap("Ricky", "Ruby"))

	subject, name := swap("Jun", "Go")
	fmt.Println(name, subject)
}
