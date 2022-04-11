package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name, attr"`
	Url    string
	Course []string
}

func main() {

	info := Website{
		Name:   "内卷大魔王【GO语言开发】",
		Url:    "www.juanjuanzi.com",
		Course: []string{"卷王的诞生——golang入门开发2022", "卷王的进阶——golang进阶开发2022", "卷王的高级——golang高级开发2022"},
	}

	file, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file) //xml的编码器
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码失败", err)
	} else {
		fmt.Println("编码成功")
	}
}
