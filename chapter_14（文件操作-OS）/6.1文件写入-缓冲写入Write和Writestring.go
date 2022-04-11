package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\黑神话悟空.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()

	str := "李子康编程大魔王13610850940\r\n"
	bufioWrite := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		bufioWrite.WriteString(str)
		bufioWrite.WriteRune('我')               //写入字符
		bufioWrite.Write([]byte{123, 124, 125}) //写入字节切片
	}
	bufioWrite.Flush()

}
