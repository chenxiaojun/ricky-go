package main

import "fmt"

type user struct {  // 定一个结构体集合类型
	name string
	age int
}

type myInt int // 定义一个自己的类型

func main() {
	var a = "init" // declare a variable
	fmt.Println(a)

	var b, c = "hello", "world"
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// short declare
	e := 23
	fmt.Println(e)

	user := user{"Ricky", 12}
	fmt.Println(user)

	var f myInt
	f = 45
	fmt.Println(f)
}