package main

import "fmt"

type Test interface {
	Test(string)
}

type Person struct {
	Name string
}

func (p *Person) Test(s string) {
	fmt.Println(p.Name)
}

func main() {
	person := &Person{"Ricky"}
	person.Test("haha")
}
