package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map 映射键到值
// map 在使用之前必须用make 而不是new来创建; 值为nil的map是空的，并且不能赋值
var m map[string]Vertex

var n = map[int]Vertex{
	2: Vertex{
		1.0, 2.9,
	},
	3: Vertex{
		4.2, 5.3,
	},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println("\n", n, n[3])
}
