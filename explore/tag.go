package main

import (
	"reflect"
	"fmt"
)

func main() {
	var u User

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}
}

type User struct {
	Name string
	Age int
}
