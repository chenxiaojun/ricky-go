package main

import "fmt"

// unlike array
func main() {
	// array
	var a [3]string
	a[0] = "haha"
	fmt.Println(a)

	var b [2][3]int
	b[0][2] = 89
	fmt.Println(b)

	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[2] = "b"
	fmt.Println(s)
	fmt.Println(len(s), cap(s)) // 3 3

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s)) // [a  b d] 4 6
	s = append(s, "e")
	fmt.Println(s, len(s), cap(s)) // [a  b d e] 5 6

}
