package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func info(greet, name string) (string, string) {
	return greet, name
}

func main() {
	c := plus(2, 5)
	fmt.Println(c)

	fmt.Println(info("Hello", "Ricky"))
	greet, name := info("HI", "DOG")
	fmt.Printf("greet: %s, name: %s\n", greet, name)

	_, name2 := info("", "Micky")
	fmt.Println(name2)
}