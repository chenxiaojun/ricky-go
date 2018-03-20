package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Hobby map[string]string
}

var myPerson Person

func newPerson() *Person {
	return new(Person)
}

func main() {
	p1 := &myPerson
	p2 := newPerson()
	p3 := &Person{}
	fmt.Printf("p1 = %v, type = %T\n", p1, p1)
	fmt.Printf("p2 = %v, type = %T\n", p2, p2)
	fmt.Printf("p3 = %v, type = %T\n", p3, p3)

	p1.Hobby = make(map[string]string)
	p1.Hobby["lang"] = "golang"
	fmt.Println(p1)
	fmt.Println(p1.Hobby)
	fmt.Println(p1.Hobby["lang"])

	fmt.Println(strings.Index("Hello World", "oW"))
}