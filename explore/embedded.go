package main

import "fmt"

type user struct {
	name string
	email string
}

// 可以组合类型
type admin struct {
	user
	level string
}


// 可以组合接口
type Aer interface {
	A() string
}

type Ber interface {
	B() int
}

type Cer interface {
	Aer
	Ber
}

type welcome string

func (w welcome) A() string {
	return "AAA"
}

func (w welcome) B() int {
	return 1024
}

func main() {
	ad := admin{ user{ "Ricky", "rickycxj@gmail.com" }, "high" }
	fmt.Println(ad.user.email) // rickycxj@gmail.com

	var c Cer
	var w welcome = "www"
	c = &w
	fmt.Println(c.A()) // AAA
	fmt.Println(c.B()) // 1024
}