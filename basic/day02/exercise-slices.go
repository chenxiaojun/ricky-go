package main

import (
	"fmt"
)

//import "code.google.com/p/go-tour/pic"
//import "tour/pic"
//
//func Pic(dx, dy int) [][]unit8 {
//	ret := make([][]unit8, dy)
//	for i := 0
//}
//
//func main() {
//	pic.Show(Pic)
//}

func main(){
	s := []int{5}  // s = 5
	s = append(s, 7) // s = [5, 7]
	s = append(s, 9) // s = [5, 7, 9]  cap 4
	x := append(s, 11) // [5, 7, 9, 11]
	y := append(s, 12)
	fmt.Printf("%p\n",&s[0])
	fmt.Printf("%p\n",&x[0])
	fmt.Printf("%p\n",&y[0])
}