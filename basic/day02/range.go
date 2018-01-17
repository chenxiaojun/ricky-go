package main

import "fmt"

var pow = []int{ 1, 2, 4, 8, 16, 32, 64, 128 }

// for循环的range格式可以对slice或者map进行迭代循环
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	test := make([]int, 10)
	for i := range test {
		test[i] = i << uint(i)
	}

	// 可以通过赋值给 _  来忽略序号和值
	// 如果只需要索引 去掉,value即可
	for _, value := range test {
		fmt.Printf("%d\n", value)
	}
}
