package main

import (
	"io"
	"log"
	"os"
)

//import "log"

//func init() {
//	log.SetPrefix("[RickyTest]")
//	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//}
//
//func main() {
//	log.Println("name：", "Ricky")
//	log.Printf("name: %s\n", "Ricky")
//}

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件关联：", err)
	}
	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Info.Println("name：", "Ricky")
	Warning.Printf("age：%d\n", 11)
	Error.Println("Ricky的测试")
}
