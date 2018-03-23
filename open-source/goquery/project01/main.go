package main

import (
	"github.com/chenxiaojun/ricky-golib/go-exercise"
	"fmt"
)

func main() {
	exercise.Init("/Users/ricky/go-2018/ricky-go/open-source/goquery/project01/config.json")
	link := "https://www.zhihu.com/people/ricky-go"
	user := exercise.NewUser(link, "")
	fmt.Println(user.GetUserID())
	fmt.Println(user.GetAvatar())
	fmt.Println(user.GetAvatarWithSize("xl"))
	fmt.Println(user.GetSign())
	fmt.Println(user.GetLocation())
}

