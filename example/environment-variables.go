package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("NAME", "Ricky")
	fmt.Println("NAME", os.Getenv("NAME"))
	fmt.Println("ERROR", os.Getenv("HA"))

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
