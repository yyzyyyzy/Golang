package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func RWFileAppend() {

	//创建文件
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\林黛玉.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//读取文件
	bufioRead := bufio.NewReader(file)
	for {
		ReadString, err := bufioRead.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(ReadString)
	}

	//写入文件
	str := "我想去字节跳动\r\n"
	bufioWrite := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		bufioWrite.WriteString(str)
	}
	bufioWrite.Flush()

}

func main() {
	RWFileAppend()
}
