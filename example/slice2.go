package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, len(s), cap(s))

	c := make([]string, len(s))
	fmt.Println(c, len(s), cap(s))
	copy(c, s)
	fmt.Println("cpy:", c, len(c), cap(c))
	fmt.Printf("%p\n", &s[0])
	fmt.Printf("%p\n", &c[0])
	c[2] = "d"
	fmt.Println(s, c)

	fmt.Println(s[1:3])
	fmt.Println(s[:2])

	d := []string{"h", "k"}
	fmt.Println(d)

	e := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		e[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			e[i][j] = i + j
		}
	}
	fmt.Println("2d: ", e)
}
