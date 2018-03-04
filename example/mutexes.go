package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops int64 = 0

	for w := 0; w < 20; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				//fmt.Println("w key:", key)
				val := rand.Intn(100)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	total := 0
	for i := 0; i < 20; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
	fmt.Println("total:", total)
}
