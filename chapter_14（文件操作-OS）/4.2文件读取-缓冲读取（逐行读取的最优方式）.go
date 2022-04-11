package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//逐行读取占用内存较小，是开发的重点
func ReadLineFile1() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\7.method的值接收者.png"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file readline fail")
	}
	defer file.Close() //最终关闭文件

	bufioRead := bufio.NewReader(file) //创建文件的读取器
	for {
		ReadLine, _, err := bufioRead.ReadLine()
		if err == io.EOF { //EOF：End Of File文件末尾
			break
		}
		fmt.Println(ReadLine)
	}
}

func ReadLineFile2() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\10.defer 1.12函数怎么执行的？"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file readline fail ")
	}
	defer file.Close()

	bufioRead := bufio.NewReader(file) //创建源文件的读取器
	for {
		ReadString, err := bufioRead.ReadString('\n') // '\n'表示按照换行符读取
		if err == io.EOF {
			break
		}
		fmt.Println(ReadString)
	}

}

func ReadLineFile3() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\8.interface前言（类型系统）？"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file readline fail ")
	}
	defer file.Close()

	bufioRead := bufio.NewReader(file)
	for {
		bytes, err := bufioRead.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(string(bytes))
	}

}

func ReadLineFile4() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\8.interface前言（类型系统）？"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file readline fail ")
	}
	defer file.Close()

	bufioRead := bufio.NewReader(file)
	for {
		ch, _, err := bufioRead.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Println(string(ch))
	}

}

func main() {
	//ReadLineFile1()
	//ReadLineFile2()
	//ReadLineFile3()
	ReadLineFile4()
}
