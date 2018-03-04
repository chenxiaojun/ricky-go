package main

import (
	"fmt"
	"sort"
)

func main() {

	// go的sort包实现了内置和用户自定义数据类型的排序功能
	strs := []string{"c", "a", "b"}
	// sort排序会直接改变传递过来的数据本身
	sort.Strings(strs)
	fmt.Println("Sorted Strings:", strs)

	ints := []int{2, 5, 1}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	fmt.Println("Int are Sorted? :", sort.IntsAreSorted(ints))
}
