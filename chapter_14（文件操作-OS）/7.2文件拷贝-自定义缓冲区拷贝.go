package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func BufferCopy(srcFileName, dstFileName string) {
	srcfile, err1 := os.Open(srcFileName)
	dstfile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err1 != nil || err2 != nil {
		fmt.Println("文件打开失败：", err1, err2)
		return
	}

	defer func() {
		srcfile.Close()
		dstfile.Close()
	}()

	bufioRead := bufio.NewReader(srcfile)
	bufioWrite := bufio.NewWriter(dstfile)

	buffer := make([]byte, 1024) //定义缓冲桶为1024字节，将会以1024字节为单位读取和写入文件
	for {
		_, err3 := bufioRead.Read(buffer)
		if err3 != nil {
			if err3 == io.EOF {
				fmt.Println("源文件读取完毕")
				break
			} else {
				fmt.Println("源文件读取错误：", err3)
				return
			}
		} else {
			_, err4 := bufioWrite.Write(buffer)
			if err4 != nil {
				fmt.Println("目标文件写入错误")
				return
			}
		}
	}

	io.Copy(bufioWrite, bufioRead)
	fmt.Println("拷贝完毕")

}

func main() {
	srcpath := "F:\\摄影软件\\pic\\丁香湖\\微信\\飞机.jpg"
	dstpath := "F:\\摄影软件\\pic\\丁香湖\\微信\\飞机(1).jpg"
	BufferCopy(srcpath, dstpath)
}
