package main

import "fmt"

// slice由函数make创建。 这会分配一个零长度的数组并且返回一个slice指向这个数组：
func main() {
	a := make([]int, 5)
	printSlice("a", a)  // len(a) = 5

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}