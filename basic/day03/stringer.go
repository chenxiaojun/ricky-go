package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type Cat struct {
	Name string
	Color string
}

func main() {
	a := Person{ "Ricky", 100 }
	fmt.Println(a)
	b := Cat { "mimi", "black" }
	fmt.Println(b)
}
