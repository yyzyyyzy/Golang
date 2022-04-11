package main

import (
	"bufio"
	"fmt"
	"os"
)

func BufferedFlush(srcFileName string) {
	byteString := []byte("我是李子康编程大魔王")
	file, _ := os.OpenFile(srcFileName, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()

	bufioWrite := bufio.NewWriter(file)
	bufioWrite.Write(byteString)
	fmt.Println(bufioWrite.Buffered(), bufioWrite.Available())

	bufioWrite.Flush() //无缓冲写入
	fmt.Println(bufioWrite.Buffered(), bufioWrite.Available())

}

func main() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\大魔王.txt"
	BufferedFlush(path)
}
