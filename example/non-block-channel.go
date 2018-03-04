package main

import (
	"fmt"
	//"time"
)

//func main() {
//	a := make(chan string)
//
//	go func(){
//		fmt.Println("abc")
//	}()
//
//	go func(){
//		a <- "ttt"
//		fmt.Println("edf")
//	}()
//
//	go func(){
//		fmt.Println("123")
//	}()
//
//	time.Sleep(time.Second*2)
//
//	fmt.Println(<-a)
//}

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case <-messages:
		fmt.Println("coming")
	default:
		fmt.Println("no messages")
	}

	a := "test"
	select {
	case messages <- a:
		fmt.Println("coming")
	default:
		fmt.Println("no messages")
	}

	select {
	case a := <-messages:
		fmt.Println(a)
	case <-signals:
		fmt.Println("coming")
	default:
		fmt.Println("no messages")
	}
}
