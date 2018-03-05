package main

import "fmt"

type user struct {
	name string
	age int
}

func main() {
	user1 := user{ "Ricky", 18 }
	fmt.Println(user1)

	user2 := user{ name: "Jack", age: 11 }
	fmt.Println(user2)

	user3 := user{ name: "Mary" }
	fmt.Println(user3)

	fmt.Println(user3.name)

	fmt.Println(&user3)
}