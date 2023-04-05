package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读取到file中，再利用ioutil将file直接读取到[]byte中（适合文件不大的情况，全部读取会占用较大内存）
// ioutil是读写工具包
func ReadFile() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\4.slice的底层数据结构？append()底层原理，扩容机制？"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("read file fail")
		return
	}
	defer file.Close()

	filebyte, err := ioutil.ReadAll(file) //filebyte是[]byte类型
	if err != nil {
		fmt.Println("read to filebyte fail")
		return
	}
	fmt.Println(string(filebyte))
}

func main() {
	ReadFile()
}
