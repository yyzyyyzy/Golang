package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func WriteBinary() {
	info := "this is my text content" // 文本内容
	file, err := os.Create("bin.bin") // 创建文件, "binbin"是文件名字
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	}
	fmt.Println("操作成功")
}

func ReadBinary() {
	file, err := os.Open("bin.bin")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var info string
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	}
	fmt.Println("解码成功")
	fmt.Println(info)
}

func main() {
	WriteBinary()
	ReadBinary()
}
