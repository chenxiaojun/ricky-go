package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", getUri)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getUri(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path) // /ricky
}
