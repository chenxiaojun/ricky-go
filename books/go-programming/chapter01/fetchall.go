package main

import (
	"strings"
	"net/http"
	"fmt"
	"time"
	"io"
	"io/ioutil"
)

func main() {
	// 运行开始的时间
	start := time.Now()

	// 初始化要读取的url数组
	target := []string { "tuchong.com", "baidu12ds.com", "deshpro.com", "weibo.com", "google.com" }

	// 初始化channel
	ch := make(chan string)

	// 遍历数组 调用fetch
	for _, url := range target {
		// 丢到线程，等待调用
		go fetch(url, ch)
	}

	// 读取线程里面的数据
	for range target {
		fmt.Println(<- ch)
	}

	//输出程序运行的总时间
	fmt.Printf("total use: %.3fs", time.Since(start).Seconds())
	fmt.Println("  ------ 使用goroutine，并发访问不同的服务器，程序执行的总时间由耗时最大的那个线程决定")
}

func fetch(url string, ch chan string) {
	// 记录开始时间
	start := time.Now()

	// 检查url前缀
	if !strings.HasPrefix(url,"https://") {
		url = "https://" + url
	}

	// 读取资源
	resp, err := http.Get(url)

	// 判断请求是否成功
	if err != nil {
		ch <- fmt.Sprintf("reading from %s: %v\n", url, err)
		return
	}

	// 将读取到的内容读取
	bytes, err := io.Copy(ioutil.Discard, resp.Body)

	// 判断请求是否成功
	if err != nil {
		ch <- fmt.Sprintf("reading from %s: %v\n", url, err)
		return
	}

	// 防止内存泄露，关闭资源
	resp.Body.Close()

	// 记录运行经过的时间
	secs := time.Since(start).Seconds()

	// 将获取的结果写入channel
	ch <- fmt.Sprintf("time=%.3fs bytes=%d url=%s status=%d", secs, bytes, url, resp.StatusCode)
}
