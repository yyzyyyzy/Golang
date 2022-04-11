package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	//创建trace文件
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//启动trace
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	//需要分析的代码
	fmt.Println("hello world")

	//停止trace
	trace.Stop()
}
