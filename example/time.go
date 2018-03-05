package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	// 当前时间
	p(time.Now())

	// 提供年月日，构建一个time
	then := time.Date(
		2018, 3, 5, 20, 20, 34, 4444, time.UTC,
	)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Local())
	p(then.Weekday())
	p(then.Before(time.Now()))
	p(then.After(time.Now()))
	p(then.Equal(time.Now()))

	diff := time.Now().Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(then.Add(diff))
	p(then.Add(-diff))
}
