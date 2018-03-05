package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	rect1 := rect{2, 4}
	fmt.Printf("area:%v\n", rect1.area())
	fmt.Printf("perim:%v\n", rect1.perim())

	rp := &rect1
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
