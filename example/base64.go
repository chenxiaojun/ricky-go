package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123#$%@@()?a=1+"

	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)

	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(string(dec))

	// URl兼容的base64
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDnc, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDnc))
}
