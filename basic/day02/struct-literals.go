package main

import "fmt"

// 结构体文法表示通过结构体字段的值 可以作为列表重新分配一个结构体
// 使用 Name: 语法可以仅列出部分字段。 (字段名的顺序无关)
// 特殊的前缀 & 返回的是一个指向结构体的指针
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // 类型为Vertex
	v2 = Vertex{ X: 1 } // Y:0 被省略
	v3 = Vertex{} // X:0 Y:0被省略
	p = &Vertex{1,2} // 类型为*Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)

	x := Vertex{3, 5}
	y := &x
	y.X = 20
	fmt.Println(x) // {20 5}
}
