package main

import "fmt"

// * 和 & 的区别
	// 1 &是地址符号，即取得某个变量的地址
	// 2 *是指针运算符，可以表示一个变量是指针类型，也可以表示一个指针变量所指向的存储单元，也就是这个地址所存储的值

	// ****指针是一个变量，其值是另一个变量的地址，即存储器位置的直接地址
type Rect struct {
	width float64
	height float64
}

func main() {
	r := Rect{100, 200}
	fmt.Printf("r=%v, address=%p\n", r, &r)
	c := r
	fmt.Printf("c=%v, address=%p\n", c, &c)
	var a *Rect = &r
	fmt.Printf("a=%v, address=%p\n", a, &a)
	b := a
	fmt.Printf("b=%v, address=%p\n", b, &b)
	fmt.Printf("*b=%v, address=%p\n", *b, &*b)

	*b = Rect{900, 200}
	fmt.Printf("r=%v, address=%p\n", r, &r)
}

//func main() {
//	r := Rect{100, 200}
//	fmt.Printf("r = %v, &r = %v\n", r, &r) // &r 可以取到r的地址
//	// * 是一个地址运算符，可以表示一个变量是指针类型
//	var p *Rect = &r
//	fmt.Println(p) // &{100 200}
//
//	// * 也可以表示指针类型 所指向的存储单元，也就是这个地址所指向的值
//	fmt.Println(*p) // {100 200}
//
//	fmt.Println(&p) // 0xc42000c030
//
//	fmt.Println(*&p) // 0xc42000c030
//
//	fmt.Println(&r)
//
//	c := &p
//	fmt.Println(c)
//
//	fmt.Printf("%p", &r)
//	fmt.Printf("%p", &r)
//	fmt.Printf("%p", c)
//	fmt.Println(*c)
//}