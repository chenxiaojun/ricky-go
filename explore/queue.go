package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// 用户 手机号
type user struct {
	id int
	name string
}

var users []user

func init() {
	// 初始化用户
	for i := 0; i < 100000; i++ {
		user := user{
			id: i,
			name: GetRandomString(4),
		}
		users = append(users, user)
	}
}

var wg sync.WaitGroup
func main() {
	start := time.Now()
	// 创建要处理到用户数对应到协程数量
	wg.Add(len(users))

	for _, user := range users {
		go sendMsg(user)
	}
	wg.Wait()
	fmt.Println("处理完毕")
	fmt.Printf("处理用时：%.2fs", time.Since(start).Seconds())
}

// 发送消息
func sendMsg(user user) {
	defer wg.Done()
	fmt.Printf("id: %d, name: %s\n", user.id, user.name)
}

// 随机用户名
// len要返回的长度
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
