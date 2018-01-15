package main

import "fmt"

func main() {
	// 在定义一个变量 但不指定其类型时(使用没有类型的var或者:=语句)， 变量的类型由右值推导得出
	// 当右值定义了类型时，新变量的类型与其相同
	v := 42 // change me
	fmt.Printf("v is of type %T\n", v)
}