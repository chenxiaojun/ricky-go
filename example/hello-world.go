package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world!")
	fmt.Fprintln(os.Stdout, "hello world!")

	fmt.Printf("%s\n", "hello world!")
	fmt.Print(fmt.Sprintf("%s\n", "hello world!"))

	fmt.Fprintf(os.Stdout, "%s\n", "hello world!")
}

// go run hello-world.go  运行go程序
// go build hello-world.go 生成二进制运行
// ./hello-world.go  执行二进制文件
