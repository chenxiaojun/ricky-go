package main

import (
	"fmt"
	"math/cmplx"
)

// Go的基本类型有
// bool string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64
// byte -> uint8的别名   rune -> int32的别名
// float23 float64 complex64 complex128
var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Println(f, ToBe, ToBe)
	fmt.Println(f, MaxInt, MaxInt)
	fmt.Println(f, z, z)


	var a int // 0
	var b float64 // 0
	var c bool // false
	var d string // ""
	fmt.Printf("%v %v %v %q\n", a, b, c, d)
}

