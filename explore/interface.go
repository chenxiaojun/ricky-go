package main

import "fmt"

type Tester interface {
	test(name string) string // go语言默认会生成一个 test(name *string) string的方法
}

type MyInt int

func (i MyInt) test(name string) string {
	return name
}

func main() {
	var t Tester                 // 申明变量t 属于接口Tester， 此时t可以是任意实现了Tester方法的参数类型变量。本质上t是一个指针，有2个字节
	var b MyInt = 6              // 设定b 属于设定的MyInt类型，其值为6
	t = b                        // 赋值给t 从而t在内存中的receiver指向b 具有MyInt里面的设定的方法
	fmt.Println(t)               // 6
	fmt.Println(t.test("Ricky")) // Ricky

	var c Tester
	c = &b
	fmt.Println(c)              // 0xc420018080
	fmt.Println(c.test("haha")) // haha
}
