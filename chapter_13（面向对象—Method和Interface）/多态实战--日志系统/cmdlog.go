package main

import (
	"fmt"
	"os"
)

type cmdWriter struct {
}

func (c *cmdWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)       //将数据序列化为字符串
	_, err := os.Stdout.Write([]byte(str)) //将数据以字节数组形式写入cmd命令行
	return err
}

func NewCmdWriter() *cmdWriter {
	return &cmdWriter{}
}
