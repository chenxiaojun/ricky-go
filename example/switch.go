package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("others")
	}

	today := time.Now().Weekday()
	fmt.Println(today) // Wednesday
	fmt.Println(time.Friday)
	fmt.Printf("%T", today)
	switch today {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thurday")
	case time.Friday:
		fmt.Println("Friday")
	default:
		fmt.Println("holiday")
	}

	fmt.Println(time.Now())

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch x := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", x)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
