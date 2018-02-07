package main

import (
	"fmt"
	"time"
)

type Person struct {
	age int
	name string
}

/** 这个属于结构体的函数，只能申明Person类型实例后才可以调用的，称之为方法 */
func (p Person) getName() string {
	return p.name
}

/** 我们一般称包级别的(直接可以通过包调用的)称之为函数 */
func GetTime() time.Time {
	return time.Now()
}

func main() {
	fmt.Println(GetTime())
	ricky := Person {
		18, "Ricky",
	}
	fmt.Println(ricky.getName())
}