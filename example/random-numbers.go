package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64() * 5) + 6)

	// 默认情况下给定的种子是确定的，每次都会产生相同的随机数数字序列
	// 要产生变化的序列，需要给定一个变化的种子
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//
	//fmt.Println(r1.Intn(100))
	//fmt.Println(r1.Intn(100))

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
