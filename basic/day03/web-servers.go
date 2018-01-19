package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
}

// 包http 通过任何实现了http.Handler的值来响应HTTP请求
func main() {
	var h Hello
	err := http.ListenAndServe("localhost: 4000", h)
	if err != nil {
		log.Fatal(err)
	}
}