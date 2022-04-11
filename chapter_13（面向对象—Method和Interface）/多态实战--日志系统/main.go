package main

import (
	"fmt"
)

func createLogger() *Logger {
	//创建日志器
	l := NewLogger()
	//创建命令行写入器
	cw := NewCmdWriter()
	//注册命令行写入器到日志器中
	l.RegisterWriter(cw)
	//创建文件写入器
	fw := NewFileWriter()
	// 设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	// 注册文件写入器到日志器中
	l.RegisterWriter(fw)
	return l
}

func main() {
	// 准备日志器
	l := createLogger()

	// 写一个日志
	l.Log("人生苦短，我爱golang")
}
