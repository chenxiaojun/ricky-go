package main

import "fmt"

func main() {
	p := new(int) // 返回的是某种类型的地址，其默认初始化值为里面变量的初始化默认值
	fmt.Println(p) // 0xc420018090
	fmt.Println(*p)

	a := a1()
	b := a1()
	fmt.Println(a == b)

	c := c1()
	d := c1()
	fmt.Println(c == d)

	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
}

func a1() *int {
	return new(int)
}

func c1() *struct{} {
	return new(struct{})
}
