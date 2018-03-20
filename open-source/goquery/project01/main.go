package main

import (
	"github.com/chenxiaojun/ricky-golib/go-exercise"
	"fmt"
)

func main() {
	exercise.Init("/Users/ricky/go-2018/ricky-go/open-source/goquery/project01/config.json")

	user := exercise.NewUser("https://www.zhihu.com/people/ricky-go", "")
	fmt.Println("user_id: ", user.GetUserID())
	//fmt.Println("location: ", user.GetLocation())
}

