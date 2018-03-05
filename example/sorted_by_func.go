package main

import (
	"fmt"
	"sort"
)

// 想自定义排序类型，需要去实现Go sort里面相关的interface，具体可参照go内部的int类型的代码

// 这里自定义一个 为内置[]string 类型的别名
type LenthSlice []string

func (p LenthSlice) Len() int           { return len(p) }
func (p LenthSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p LenthSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var animals LenthSlice
	animals = []string{"Dog", "People", "Tiger"}
	fmt.Println(animals.Len())

	sort.Sort(animals)
	fmt.Println(animals)
}
