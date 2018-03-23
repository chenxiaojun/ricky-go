package main

import (
	"fmt"
	"time"
	"runtime"
	"net/http"
)

func main() {
	start := time.Now()
	http.HandleFunc("/test", do)
	http.ListenAndServe(":8000", nil)
	fmt.Println("\ntotal:", time.Since(start).Nanoseconds())
}

func do(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	runtime.GOMAXPROCS(4)
	for i := 0; i < 1000; i++ {
		go add(i, ch)
		<-ch
	}
}

func add(end int, ch chan int) {
	sum := 0
	for i := 0; i < end; i++ {
		sum += i
	}
	ch <-sum
}
