package main

import "fmt"

// Go 具有指针。 指针保存了变量的内存地址
func main() {
	i, j := 42, 2701

	p := &i //point to i
	fmt.Println(*p) //read i through the pointer

	*p = 21 // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j

	//p is a mac address, *p is the value of this address
	fmt.Println(p, *p) //0xc420018088 73
}
