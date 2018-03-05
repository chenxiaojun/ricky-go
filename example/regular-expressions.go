package main

import (
	"fmt"
	"regexp"
)

// 正则表达式
func main() {
	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面是直接使用字符串， 但是对于一些其它但正则任务，需要Compile一个优化但Regexp结构体
	r, err := regexp.Compile("p([a-z]+)ch")
	fmt.Println(err)
	fmt.Println(r.MatchString("peach"))

	// 查找匹配字符串的
	fmt.Println(r.FindString("peach punch"))

	// 返回匹配到的开始位置和结束位置
	fmt.Println(r.FindStringIndex("peach punch"))

	// 返回完全匹配和子模式匹配到到内容
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 带All的会返回所有的匹配项
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 上面用了MatchString这样的方法，我们可以提供[]byte参数把String从函数中剔除
	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.Find([]byte("peach punch")))

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
}
