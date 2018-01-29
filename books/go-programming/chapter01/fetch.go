package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	target := []string{ "https://tuchong.com/explore", "http://ricky.xin" }
	for _, url := range target {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
