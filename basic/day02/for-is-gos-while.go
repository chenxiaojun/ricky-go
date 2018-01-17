package main

import "fmt"

// 其它语言中的while 对应到Go中就是for
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
