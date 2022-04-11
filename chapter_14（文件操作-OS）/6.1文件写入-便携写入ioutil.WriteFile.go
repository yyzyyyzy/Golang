package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func EasyWriteFile() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\鲁智深.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		file.Close()
	}()

	str := "李子康编程大魔王13610850940"

	err1 := ioutil.WriteFile(path, []byte(str), 0666)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("write success")
}

func main() {
	EasyWriteFile()
}
