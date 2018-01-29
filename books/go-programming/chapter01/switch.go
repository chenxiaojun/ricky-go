package main

import "fmt"

//func main() {
//	a := "Ricky"
//
//	switch a {
//	case "Amy":
//		fmt.Print("Amy")
//	case "Ricky":
//		fmt.Print("Ricky")
//	}
//}

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a) // 1
	fmt.Println("&a = ",&a) // 0xc420018090
	fmt.Println("*&a = ",*&a) // 1
	fmt.Println("b = ",b) // 0xc420018090
	fmt.Println("&b = ",&b) // 0xc42000c028
	fmt.Println("*&b = ",*&b) // 0xc420018090
	fmt.Println("*b = ",*b) // 1
	fmt.Println("c = ",c) // 0xc42000c028
	fmt.Println("*c = ",*c) // 0xc420018090
	fmt.Println("&c = ",&c) // 0xc42000c030
	fmt.Println("*&c = ",*&c) // 0xc42000c028
	fmt.Println("**c = ",**c) // 1
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c) // 1
	fmt.Println("x = ",x) //  1

	test(&*b)
}

func test(v *int) { // v接收的必须是一个地址变量
	fmt.Println(v)
	fmt.Println(*v) // 1
}