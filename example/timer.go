package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("go")

	// 用来阻塞程序向下执行，知道上面的定时器失效
	<- timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	fmt.Println("go 2")
	go func(){
		<- timer2.C
		fmt.Println("timer 2 expired")
	}()
	// 取消掉定时器
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}
	// 如果仅仅是想要单纯的等待，那么使用time.Sleep()就好了
	// 之所以使用定时器，是因为定时器可以随时取消掉
}