package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\11.panic和recover函数"
	file1, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	} else {
		fmt.Println("文件打开成功")
	}

	defer func() {
		file1.Close()
		fmt.Println("文件已关闭")
	}()

	fmt.Println("我想打开文件5s")
	time.Sleep(5 * time.Second)
}
