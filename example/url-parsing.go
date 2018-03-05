package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "http://www.baidu.com:80?key=golang&sort=desc"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Host)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Query())
	fmt.Println(url.ParseQuery(u.RawQuery))

	h := strings.Split(u.Host, ":")
	fmt.Println(h[1])
}
