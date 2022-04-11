package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func bufferCopyFile(srcFilePath, dstFilePath string) (writeLength int64, err error) {
	srcfile, err1 := os.Open(srcFilePath)
	dstFile, err2 := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		return
	}
	defer func() {
		srcfile.Close()
		dstFile.Close()
	}()

	bufioRead := bufio.NewReader(srcfile)
	bufioWrite := bufio.NewWriter(dstFile)

	return io.Copy(bufioWrite, bufioRead)
}

func main() {
	srcfile := "F:\\摄影软件\\pic\\新年\\DSC_0803.MOV"
	dstfile := "F:\\摄影软件\\pic\\新年\\DSC_0803(1).MOV"
	writelength, err := bufferCopyFile(srcfile, dstfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(writelength)

}
