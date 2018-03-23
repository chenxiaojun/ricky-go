package main

import (
	//"fmt"
	"net/http"
	//"os"
)

var total int

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		totalPlus()
		//fmt.Fprintf(os.Stdout, fmt.Sprintf("total: %d\n", total))
	})
	http.ListenAndServe(":8000", nil)
}

func totalPlus(){
	total = total + 1
}
