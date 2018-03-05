package main

import (
	"fmt"
	"os"
)

// defer 常用来 程序执行结束后的清尾工作，如其它语言执行完毕会回调一样

func main() {
	f := createFile("/tmp/t.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
