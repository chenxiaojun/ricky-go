package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 30
	fmt.Println(n)
	fmt.Println(math.Sin(n))
}
