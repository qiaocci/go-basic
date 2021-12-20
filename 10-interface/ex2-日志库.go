package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger interface {
	Info(string)
}

type ConsoleLogger struct {
}

func (c *ConsoleLogger) Info(x string) {
	fmt.Printf("[%v] %v\n", time.Now().Format("2006-01-02 15:04:05"), x)
}

type FileLogger struct {
	filename string
}

func (fl *FileLogger) Info(x string) {
	var f *os.File
	var err1 error
	if checkFileExists(fl.filename) {
		f, err1 = os.OpenFile(fl.filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件

		if err1 != nil {
			fmt.Printf("打开文件失败%v", err1)
			return
		}
	} else {
		f, err1 = os.Create(fl.filename)
		if err1 != nil {
			fmt.Printf("创建文件失败%v", err1)
			return
		}
	}

	defer f.Close()

	n, err1 := io.WriteString(f, x+"\n")
	if err1 != nil {
		fmt.Printf("err=%#v", err1)
		panic("写入失败!")
	}
	fmt.Printf("写入%v个字节\n", n)
}

func checkFileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	var logger Logger

	c := ConsoleLogger{}
	c.Info("ConsoleLogger打印日志")

	//println(checkFileExists("xx.go"))
	f := &FileLogger{"测试.txt"}
	logger = f
	logger.Info("你")
}
