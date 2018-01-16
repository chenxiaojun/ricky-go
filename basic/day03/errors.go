package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error{
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
// Go程序用error值来表示错误状态 与fmt.Stringer类似， `error`类型是一个内建接口
// 通常函数会返回一个error值，调用它的代码应当判断这个错误是否等于nil，来进行错误处理
func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}