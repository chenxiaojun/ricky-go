// & 和 * 的区别
	// 每个变量 都占用一个内存单元，而这个内存单元有自己都地址，访问的时候可以根据这个地址找到它
	// 指针也是一个变量，但是这个变量的值对应的是另外一个变量的地址，所以指针也有自己的地址单元
package main

import "fmt"

// 申明一个变量test，它是一个map类型
var test map[string]int = map[string]int{"a": 10, "b": 11}

func main(){
	// 申明一个指针变量，这个指针指向一个map类型的地址
	var p *map[string]int
	p = &test

	fmt.Printf("test=%v, test address = %p\n", test, &test) // test=map[a:10 b:11], test address = 0x1140040
	fmt.Printf("p=%v, p address = %p\n", p, &p) // p=&map[a:10 b:11], p address = 0xc42000c028

	// 将变量赋值给另外一个变量, 采用后进先出的形式添加到新到变量中
	test2 := test
	fmt.Printf("test2=%v, test2 address = %p\n", test2, &test2) // test2=map[a:10 b:11], test2 address = 0xc42000c038

	// 将指针赋值给一个新到指针，和变量的赋值是类似的，也会开一个新的地址存储这个指针
	p2 := p
	fmt.Printf("p2=%v, p2 address = %p\n", p2, &p2) // p2=&map[a:10 b:11], p2 address = 0xc42000c040

	// 对于指针变量，可以使用*p的形式 来访问到它所指向到变量的值
	fmt.Printf("*p2=%v, *p2 address = %p\n", *p2, *p2) // *p2=map[a:10 b:11], *p2 address = 0xc420072180
}
