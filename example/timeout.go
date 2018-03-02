package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second * 2)
		c1 <- "c1 coming"
	}()
	select{
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	} // timeout


	c2 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "c2 coming"
	}()
	select{
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	} // c2 coming
}

