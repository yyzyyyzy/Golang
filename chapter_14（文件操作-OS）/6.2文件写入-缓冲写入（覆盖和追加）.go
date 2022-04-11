package main

import (
	"bufio"
	"fmt"
	"os"
)

func CoverWrite() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\贾宝玉.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "李子康\r\n"
	bufioWrite := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		bufioWrite.WriteString(str)
	}
	bufioWrite.Flush()
}

func AppendWrite() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\贾宝玉.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()

	str := "林黛玉\r\n"
	bufioWrite := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		bufioWrite.WriteString(str)
	}
	bufioWrite.Flush()
}

func main() {
	CoverWrite()
	AppendWrite()
}
