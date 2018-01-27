package chapter01

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	fmt.Println(counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	box := []string{}
	for input.Scan() {
		counts[input.Text()] ++
		// 输出出现重复行的文件名称
		if counts[input.Text()] >= 2 {
			if !checkExists(f.Name(), box) {
				box = append(box, f.Name())
			}
		}
	}
	fmt.Println(box)
}

func checkExists(file string, box []string) bool {
	flag := false
	for _, val := range box {
		if file == val {
			flag = true
			break
		}
	}
	return flag
}