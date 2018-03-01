package main

import "fmt"

func main() {
	i := 1
	var s int
	for i <= 10 {
		s += i
		i ++
	}
	fmt.Println(s) // 55

	for j := 5; j < 6; j++ {
		s += j
	}
	fmt.Println(s) // 60

	for {
		fmt.Println("loop")
		break
	}

	for k := 0; k < 10; k++ {
		if k % 2 == 0 {
			continue
		}
		fmt.Printf("%v ", k) // 1 3 5 7 9
	}
}
