package main

import (
	"fmt"
	"strings"
)

// 返回目标字符串t出现的第一个索引位置，或者在没有的时候返回-1
func Index(obj []string, target string) int {
	for i, v := range obj {
		if v == target {
			return i
		}
	}
	return -1
}

// 如果目标字符串t在这个切片中，则返回true
func Include(obj []string, target string) bool {
	return Index(obj, target) >= 0
}

// 如果这些切片中的字符有一个满足条件f 则返回true
func Any(obj []string, f func(string) bool) bool {
	for _, v := range obj {
		if f(v) {
			return true
		}
	}
	return false
}

// 如果切片中所有字符都满足条件f 则返回true
func All(obj []string, f func(string) bool) bool {
	for _, v := range obj {
		if !f(v) {
			return false
		}
	}
	return true
}

// 返回一个新的切片，新的切片返回包含所有满足条件的字符
func Filter(obj []string, f func(string) bool) []string {
	newSlice := make([]string, 0)
	for _, v := range obj {
		if f(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// 返回一个对原始切片中所有字符串执行函数f后对新切片
func Map(obj []string, f func(string) string) []string {
	newSlice := make([]string, len(obj))
	for i, v := range obj {
		newSlice[i] += f(v)
	}
	return newSlice
}

func main() {
	strs := []string{"A", "O", "I", "O", "U"}
	fmt.Println(Index(strs, "O"))
	fmt.Println(Include(strs, "O"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "O")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "O")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "O")
	}))

	fmt.Println(Map(strs, strings.ToLower))
}
