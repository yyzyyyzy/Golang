package main

import (
	"fmt"
	"io"
	"os"
)

func SeekFunc(filename string) {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()

	buffer := make([]byte, 1)
	file.Read(buffer)
	fmt.Println(string(buffer))

	file.Seek(4, io.SeekStart) //光标的位置相对于文件开头
	file.Read(buffer)
	fmt.Println(string(buffer))

	file.Seek(4, io.SeekCurrent) //光标的位置相对于当前位置
	file.Read(buffer)
	fmt.Println(string(buffer))

	file.Seek(4, io.SeekEnd) //光标的位置相对于文件末尾
	file.Read(buffer)
	fmt.Println(string(buffer))
}

func main() {
	path := "E:\\golandlearning\\chapter_14（文件操作-OS）\\葫芦娃.txt"
	SeekFunc(path)
}
