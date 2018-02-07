package main

import (
	"sync"
	"io"
	"errors"
	"log"
)

// 定义一个资源池已经关闭的错误
var ErrPoolClosed = errors.New("资源池已经关闭。")

// 一个安全的资源池， 被管理的资源必须都实现io.Close接口
type Pool struct {
	m sync.Mutex // 互斥锁
	res chan io.Closer // 有缓冲的通道，用于保存共享的资源
	factory func() (io.Closer, error) // 用于创建一个新的资源
	closed bool // 用于判断资源池是否被关闭
}

/**
	fn 创建新资源的函数，size指定资源池的大小
 */
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	return &Pool{
		factory: fn,
		res: make(chan io.Closer, size),
	}, nil
}

// 从资源池里获取资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire: 共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: 新生成资源")
		return p.factory()
	}
}
































