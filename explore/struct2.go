package main

import "fmt"

type User struct {
	Page *Page
	Name string
}

type Page struct {
	Link string
}

func (page *Page) getLink(link string) {
	fmt.Println(link)
}

func main() {
	user := User{
		&Page{"http://www.baidu.com"},
		"Ricky",
	}

	user.Page.getLink("http://www.ttt.com")
}