package chapter01

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 获取命令行参数三种方式 及 简单执行时间评估
func main() {
	start_one := time.Now()
	for key, val := range os.Args[1:] {
		fmt.Println(key, val)
	}
	fmt.Printf("%.2fs cost\n", time.Since(start_one).Seconds())

	start_two := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs cost\n", time.Since(start_two).Seconds())

	start_three := time.Now()
	var box string
	for i := 1; i < len(os.Args); i++ {
		box += os.Args[i]
		box += " "
	}
	fmt.Println(box)
	fmt.Printf("%.2fs cost\n", time.Since(start_three).Seconds())
}
