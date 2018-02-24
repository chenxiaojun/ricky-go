package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var u User
	h := `{"name":"张三","age":15}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}

	// 通过Marshal可以将结构体转化为json对象， 通过Unmarshal可以将json对象转化为结构体
	newJson, err := json.Marshal(&u)
	fmt.Println(string(newJson))
}

type User struct {
	Name string
	Age  int
}
