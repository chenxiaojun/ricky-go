package main

import (
	"fmt"
	"regexp"
)

func main() {
	pat := `.*(.gif)`
	obj := `a.png picture.gif`

	// MatchString 判断字符串能否在正则表达式中被匹配到
	fmt.Println(regexp.MatchString(pat, obj)) // true, nil

	// QuoteMeta(s string) string 将s中的正则表达式元字符转义成普通字符
	fmt.Println(regexp.QuoteMeta(obj))

	// Compile
	// 将正则表达式编译成一个正则对象(使用PERL语法)
	// 该正则对象会采用leftmost-first模式，选择第一个匹配到的结果
	// 如果正则表达式语法错误，则返回错误信息

	// MustCompile 同Compile，但是在解析失败的时候 会产生panic
	email := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	fmt.Println(email.MatchString("rickycxj@gmail.com"))
	fmt.Println(email.MatchString("ricky com"))
	fmt.Println(email.String())
	fmt.Println(email.FindString("rickycxj@gmail.com"))
}