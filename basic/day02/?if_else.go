// ? 这个里面输出的为什么不是 9  27 >= 27 20 反而输出的是 27 >= 27 9 20
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不可以使用v了
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		"---",
		pow(3, 3, 20),
	)

	fmt.Println(4, 8)
}
