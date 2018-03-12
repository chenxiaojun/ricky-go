package main

import "fmt"

// interface 是一组method的组合。
// 空interface(interface{})不包含任何的method，因此所有的类型都实现了空interface
// 空interface在我们需要存储任意类型数值的时候相当有用 有点类似C中的void *

func main() {
	slice := make([]interface{}, 10)
	map1 := make(map[string]string)
	map2 := make(map[string]int)
	map2["TaskID"] = 1
	map1["Command"] = "ping"
	map3 := make(map[string]map[string]string)
	map3["mapvalue"] = map1
	slice[0] = map2
	slice[1] = map1
	slice[3] = map3
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[3])

	who := Greeting("Ricky", "Micky")
	for _, people := range who {
		fmt.Println(people)
	}
}

// Go语言函数中 ...表示为可变参数，可以接受任意个数的参数
//  返回一个slice
func Greeting(who ...string) []string {
	fmt.Println(who)
	fmt.Printf("%T\n", who) // []string
	return who
}