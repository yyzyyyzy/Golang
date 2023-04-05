package main

import (
	"errors"
	"fmt"
	"os"
)

type fileWriter struct {
	file *os.File
}

// 设置文件写入器 写入的文件名
func (f *fileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	return err
}

func (f *fileWriter) Write(data interface{}) error {
	if f.file != nil {
		return errors.New("创建文件失败") //日志文件还没有创建
	}
	str := fmt.Sprintf("%v\n", data)    //将数据序列化为字符串
	_, err := f.file.Write([]byte(str)) //将数据以字节数组写入文件中
	return err
}

// 创建文件写入器实例
func NewFileWriter() *fileWriter {
	return &fileWriter{}
}
