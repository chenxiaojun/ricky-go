package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 从字符串中解析数字在很多程序中是一个基础常见的任务
	f, _ := strconv.ParseFloat("1.23", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
}
