package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 3)
	sum(2, 6, 7)

	nums := []int{2, 6, 3, 1}
	sum(nums...)
}
