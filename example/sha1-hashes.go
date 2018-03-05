package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	m5 := md5.New()
	m5.Write([]byte(s))
	ms := m5.Sum(nil)
	fmt.Printf("%x\n", ms)
	fmt.Printf("%x\n", m5.Sum([]byte("b")))
}
