package main

import "fmt"

// 一个结构体(`struct`)就是一个字段的集合
// type的含义跟其字面的意思相符
type Vertex struct{
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2 })

	// 结构体字段使用点号来访问
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v) //4 {4, 2}

	// 结构体字段可以使用结构体指针来访问
	// 通过指针间的访问是透明的
	p := &v

	p.X = 1e9
	fmt.Println(v) // {1000000000, 2}
}