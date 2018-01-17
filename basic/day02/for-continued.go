package main

import "fmt"

// 可以让前置、后置语句为空
func main() {
	sum := 1
	for ; sum < 500; {
		sum += sum
	}
	fmt.Println(sum)
}
