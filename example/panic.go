package main

import (
	"fmt"
	"os"
)

func main() {

	// panic 意味着有些出乎意料的错误发生。 通常我们用它来表示程序正常运行中不该出现的
	// 或者我们没有处理好的数据
	//panic("a problem")

	// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道(或者不想)处理时终止运行
	// 这里是一个在创建一个新文件时返回异常错误时的panic用法
	_, err := os.Create("/place/file")
	if err != nil {
		panic(err) // panic: open /place/file: no such file or directory
	}

	fmt.Println("ending")
}
