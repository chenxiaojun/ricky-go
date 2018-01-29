package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(),	*sep)) // 默认连接符为 " "， 如果 -s -- 则默认为 --
	fmt.Println(*n)
	fmt.Println(*n == false)
	var color []string
	if color == nil { // color == nil 这种可以判断  color == false 这种不可以判断
		fmt.Println("haha")
	}
	if !*n {
		fmt.Println() // 默认去掉末尾的换行符 如果传了-n 就不去掉了，因为此时*n为true
	}
}
