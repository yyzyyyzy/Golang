package main

import (
	"fmt"
	"os"
)

func CheckFileIsNotExist(srcFileName string) {
	_, err := os.Stat(srcFileName)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Println("文件存在，绝对路径为：", srcFileName)
	}
}

func main() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\6.struct和内存对齐"
	CheckFileIsNotExist(path)
}
