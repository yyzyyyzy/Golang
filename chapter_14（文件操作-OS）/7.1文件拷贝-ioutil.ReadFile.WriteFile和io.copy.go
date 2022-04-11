package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFileFunction1(srcFileName string, dstFileName string) {
	srcFile, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(dstFileName, srcFile, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝成功")
}

func CopyFileFunction2(srcFileName string, dstFileName string) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
}

func main() {
	srcfile := "E:\\golandlearning\\chapter_14（文件操作-OS）\\林黛玉.txt"
	dstfile := "E:\\golandlearning\\chapter_14（文件操作-OS）\\林黛玉(1).txt"
	CopyFileFunction1(srcfile, dstfile)
	CopyFileFunction2(srcfile, dstfile)
}
