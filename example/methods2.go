package main

import "fmt"

type T struct {
	val int
}

func (t T) m1() {
	t.val += 10
}

func (t *T) m2() {
	t.val += 10
}

func main() {
	t := T{ val: 30 }
	fmt.Println(t) // 30
	t.m1()
	fmt.Println(t) // 30

	t2 := T{ val: 30 }
	fmt.Println(t2) // 30
	t2.m2()
	fmt.Println(t2) // 40
}
