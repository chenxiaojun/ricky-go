package main

import (
	"fmt"
	"encoding/json"
)

type Account struct {
	Account string
	Password string
}

func main() {
	var s Account
	str := `{"account": "rickycxj@gmail.com", "password": "test123"}`
	fmt.Println([]byte(str))
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Account)
}