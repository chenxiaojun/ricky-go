package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// %+v的格式化输出内容包括结构体的字段名
	fmt.Printf("%+v\n", p)

	// %#v的格式化输出这个的 运行源代码片段
	fmt.Printf("%#v\n", p)

	// %T 查看变量的类型
	fmt.Printf("%T\n", p)

	// %t 格式化bool类型
	fmt.Printf("%t\n", true)

	// %d 数字类型
	fmt.Printf("%d\n", 123)

	// %b 二进制类型
	fmt.Printf("%b\n", 10)

	// %c 字符类型
	fmt.Printf("%c\n", 33)

	// %x 提供十六进制编码
	fmt.Printf("%x\n", 456)

	// %s 字符串类型
	fmt.Printf("%s\n", "haha")

	// 像Go源代码那样有双引号输出
	fmt.Printf("%q\n", "\"string'ha")

	// 浮点数类型格式化
	fmt.Printf("%.1f\n", 77.77)

	// 输出一个指针的值
	fmt.Printf("%p\n", &p)

	// 控制输出结果的宽度和精度 可以在%后面使用数字来控制输出宽度
	fmt.Printf("|%6d||%-6d|\n", 12, 34)

	// Sprintf 格式化后不输出，返回的是格式化好的字符串
	fmt.Println(fmt.Sprintf("a %s", "string"))

	// Fprintf来格式化并输出到io.Writers而不是os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
