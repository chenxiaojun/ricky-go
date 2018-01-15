package main

import (
	"fmt"
	"strings"
)

func WorldCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range(strings.Fields(s)) {
		m[w] ++
	}
	return m
}

func main() {
	str := "Hello This is my best friend Ricky"
	fmt.Println(WorldCount(str))
}
