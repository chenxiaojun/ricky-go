package main

import "fmt"
import s "strings" // 别名

// 给fmt.Println 一个短的别名
var p = fmt.Println

// 字符串处理函数
func main() {
	// 查看是否包含
	p("Contains: ", s.Contains("ricky", "c"))

	// 包含的次数统计
	p("Counts: ", s.Count("ricky chan", "c"))

	// 是否包含前缀
	p("Prefix: ", s.HasPrefix("ricky chan", "ri"))
	p("Suffix: ", s.HasSuffix("ricky chan", "ri"))

	str := "ricky"
	p(str[0:len("k")] == "r")

	// 返回字符串所在位置的索引
	p("Index: ", s.Index("ricky chan", "c"))

	// 字符串连接
	p("Join: ", s.Join([]string{"a", "b", "c"}, "-"))

	// 字符串重复，Repeat
	p("Repeat: ", s.Repeat("a", 5))

	// 字符串替换 replace
	p("Replace: ", s.Replace("ccn", "c", "m", -1))
	p("Replace: ", s.Replace("ccn", "c", "m", 1))

	// 字符串分割
	p("Split: ", s.Split("a-b-c-d-e", "-"))

	// 大小写转换
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
}
