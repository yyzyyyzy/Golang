package main

import (
	"fmt"
	"os"
)

func Openfile() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\反恐精英.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()

}

func Createfile() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\冒险岛.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
}

func main() {
	Createfile()
	Openfile()
}
