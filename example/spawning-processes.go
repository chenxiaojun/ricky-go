package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// exec.Command 函数帮助我们 创建一个表示这个外部进程的对象
	dateCmd := exec.Command("date")

	// .Output是另一个帮助我们处理运行一个命令的常见情况的函数，它等待命令运行完成
	// 并收集命令的输出。 如果没有错，dateOut将获取到日期信息的字节
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	//lsCmd := exec.Command(`ls`, "-al")
	//
	//lsOut, err := lsCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("> ls -l")
	//fmt.Println(string(lsOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()

	grepIn.Write([]byte("hello grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 如果不想像上面那样，一行命令拆分成多个参数，可以使用bash来执行
	lsCmd := exec.Command("bash", "-c", "ls -alh")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("ls -alh")
	fmt.Println(string(lsOut))
}
