package main

import (
	"database/sql"
	sysfmt "fmt" // Go默认查找包的顺序 GOROOT > GOPATH > 远程库
	_ "net/http" // 有时候我们需要导入一个包，但是又不使用它，按照规则，这是不可行的。 但是可以在其前面加上_重命名包后就可以了
)

import _ "github.com/go-sql-driver/mysql" // 我们仅仅是想执行这个mysql包里面的init方法，并不想使用这个包

func init() {
	sysfmt.Println("coming") // init优先main执行，一个包可以有多个init
}

func init() {
	sysfmt.Println("coming...")
}

func main() {
	sysfmt.Println("测试") // 给包取一个别名

	db, _ := sql.Open("mysql", "root:@/rtest?charset=utf8")
	rows, _ := db.Query("SELECT `name` FROM a")

	for rows.Next() {
		var name string
		rows.Columns()
		rows.Scan(&name)
		sysfmt.Println(name)
	}
}
