// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"strconv"
	"math/rand"
	"sync"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{}) // 用户管理线程
	rand.Seed(time.Now().UTC().UnixNano()) // 随机数种子
	var mu sync.Mutex // 锁机制

	go func() { // 开启线程
		defer c.Close() // 一个线程执行完 自动将其关闭
		defer close(done)
		start, _ := strconv.Atoi(os.Args[1])
		base := start
		end, _ := strconv.Atoi(os.Args[2])
		var max = 0
		var max_user_id = 0

		for start < end {
			mu.Lock()
			_, _, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			given_price := rand.Intn(9999)
			if given_price > max {
				max = given_price
				max_user_id = start
			}
			number := start - base
			log.Printf("[user: user_%d] [connect] 最高出价:user_%d(%d), 出价人数：%d", start, max_user_id, max, number)
			start++
			mu.Unlock()
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}