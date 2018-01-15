package main

import "fmt"

// Go的返回值可以被命名，并且像变量那样使用
// 返回值的名称应当具有一定的意义，可以作为文档使用
// 没有参数的return语句返回结果的当前值。 也就是直接返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
