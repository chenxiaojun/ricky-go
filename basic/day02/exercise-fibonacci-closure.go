package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	n := 0

	return func() int {
		if n == 0 {
			n++;
			return a; // f(a) = f(a - 1) + f(a - 2)  f(1) = 0
		}else if n == 1{
			n++; // f(2) = 1
			return b;
		}
		a, b = b, a + b
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

