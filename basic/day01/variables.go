package main

import "fmt"

// var语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面
// var可以定义在包或函数的级别
var c, python, java bool

// var 定义变量的时候可以初始化一些值
var x, y int = 2, 6

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// 如果变量初始化时候是表达式，可以省略类型申明，因为这时候可以从变量里面获取到它所属的类型
	var name, age = "Ricky", 100
	fmt.Println(x, y, name, age)

	// 在函数中，`:=` 简洁的赋值语句在明确类型的地方，可以用于替代var定义
	// 函数外每个语句都必须以关键字开始(`var`, `func`等), `:=`结构不可以用用于函数外面
	k := 3
	fmt.Println(k)
}