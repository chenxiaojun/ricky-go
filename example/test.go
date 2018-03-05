package main

import (
	"fmt"
)

type test struct {
	name string
}

func main() {
	ops := make(chan test)
	test1 := test{"dog"}
	go func() { ops <- test1 }()
	//fmt.Println(<-ops)
	for r := range ops {
		fmt.Println(r.name)
		close(ops)
	}
}
