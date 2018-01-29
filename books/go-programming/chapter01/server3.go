package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost: 8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 打印http请求头
	for key, val := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, val)
	}

	// 打印http请求方式，请求的地址，请求的协议
	fmt.Fprintf(w, "Method=%q, Url=%q, Proto=%q\n", r.Method, r.URL.Path, r.Proto)

	// 打印请求来的主机，请求的ip地址
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)

	// 得到表单数据
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// 遍历表单
	for key, val := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, val)
	}
}