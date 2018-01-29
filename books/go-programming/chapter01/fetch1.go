package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
)

func main() {
	urls := []string { "tuchong.com/events/" }

	for _, url := range urls {
		// 检查是否带有https前缀, 如果没有自动添加
		if !strings.HasPrefix(url,"https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: error %v\n", err)
			os.Exit(1)
		}

		size, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 关闭Body数据流，防止资源泄露

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: error %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("the size is %d\n", size)
		fmt.Printf("the status is %d", resp.StatusCode) //输出http状态码
	}
}
