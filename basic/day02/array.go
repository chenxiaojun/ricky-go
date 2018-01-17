package main

import "fmt"

// [n]T 是一个有n个类型为T的值的数组
// 表达式
// var a [10]int
// 数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心。 Go提供来更加便利的方式来使用数组
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // Hello World
	fmt.Println(a) // [Hello World]

	// 一个slice 会指向一个序列的值，并且包含来长度信息
	p := []int{ 2, 3, 5, 7, 11, 13 }
	fmt.Println("p ==", p) // [2 3 5 7 11 13]

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}
