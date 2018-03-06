package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Go提供需要执行的可执行文件的绝对路径，使用exec.LookPath来得到它
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// 用作exec的参数
	// exec的参数是切片形式的(不是放在一起的一大串)
	// 我们使用exec.LookPath得到它
	args := []string{"ls", "-a", "-l", "-h"}

	// exec需要使用到环境变量，这里提供当前的环境变量
	env := os.Environ()
	fmt.Println(env)

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
