package main

import (
	"fmt"
	"io/ioutil"
)

func ReadDirfunction() {
	files, _ := ioutil.ReadDir("chapter_5（工程管理）")
	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			fmt.Println("文件夹")
		} else {
			fmt.Println("文件")
		}
	}
}

func main() {
	ReadDirfunction()
}
