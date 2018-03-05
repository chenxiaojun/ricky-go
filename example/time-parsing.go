package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	t := time.Now()
	p(t.Format(time.RFC3339))
	p(t.Format("3.04PM"))
	p(t.Format("Mon Jan 2 15:04:05 2006"))

	form := "3 04 PM"
	t2, _:= time.Parse(form, "8 41 PM")
	p(t2)
}