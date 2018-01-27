package chapter01

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	container := ""
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		container += input.Text()
		container += " "
		counts[input.Text()]++
		// 如果有重复3次的就直接退出
		if counts[input.Text()] >= 3 {
			break
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

	fmt.Println(container)
	fmt.Println(counts)
}
