package main

import "fmt"

func f(flag string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("%s: %d\n", flag, i)
	}
}

func main() {
	f("A")

	go f("B")

	go func(flag string){
		for i := 0; i < 2; i++ {
			fmt.Printf("%s: %d\n", flag, i)
		}
	}("C")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}