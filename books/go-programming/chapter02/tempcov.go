package main

import "fmt"

type Celsius float64 // 申明摄氏温度类型
type Fahrenheit float64 // 申明华氏温度类型，虽然它跟Celsius一样的底层float64，但是属于不同的类型，不可比较

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC = 0
	BoilingC = 100
)

func main() {
	fmt.Println(CToF(20))
	fmt.Println(FToC(20))
	fmt.Println(BoilingC)

	var name string = "Ricky"
	fmt.Println(name)

	c := AbsoluteZeroC
	c.asd()

	var f Fahrenheit = 3.12
	fmt.Println(f.aaa())
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(f - 32) * 5/ 9
}

func (c Celsius) asd() string {
	return fmt.Sprintf("is is %g\n", c)
}

func (f Fahrenheit) aaa() string {
	return fmt.Sprintf("haha f = %g\n", f)
}
