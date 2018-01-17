package main

import "fmt"

// slice的零值是`nil` 一个nil的slice长度和容量是0
func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
