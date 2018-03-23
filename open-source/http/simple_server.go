package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/ricky", IndexHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}
