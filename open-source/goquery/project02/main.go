package main

import (
	"github.com/DeanThompson/zhihu-go"
	"fmt"
)

func main(){
	zhihu.Init("/Users/ricky/go-2018/ricky-go/open-source/goquery/project02/config.json")

	user := zhihu.NewUser("https://www.zhihu.com/people/ricky-go", "")
	fmt.Println("..", user.GetLocation())
}


