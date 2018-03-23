package main

import (
	"net/http"
	"fmt"
	"os"
)

var total_money = 10000

func main() {
	http.HandleFunc("/test", do)
	http.ListenAndServe(":8000", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	go take_money(10)
	fmt.Fprintf(os.Stdout, fmt.Sprintf("total: %d\n", total_money))
}

func take_money(i int) int {
	total_money = total_money - i
	return total_money
}
