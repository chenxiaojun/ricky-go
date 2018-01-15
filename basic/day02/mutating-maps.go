package main

import "fmt"

// 在map m中插入或修改一个元素 m[key] = elem
// 获得元素 elem = m[key]
// 删除元素 delete(m, key)
// 通过双赋值检测某个键存在 elem, ok = m[key]
// 如果key在m中， `ok`为true。 否则，ok为`false`
// elem是map的元素类型的零值，同样的，当从map中读取某个不存在的键时，结果是map的元素类型的零值
func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok, m["Answer23∂"])

	m["age"] = 122
	a, b := m["age"]
	fmt.Println("The value:", a, "Present?", b, m["age"])
}