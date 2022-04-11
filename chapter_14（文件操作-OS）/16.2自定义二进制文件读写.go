package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website1 struct {
	Url int32
}

func WriteUserBinary() {
	file, err := os.Create("output(1).bin")
	for i := 1; i <= 10; i++ {
		info := Website1{int32(i)}
		if err != nil {
			fmt.Println("文件创建失败 ", err.Error())
			return
		}
		defer file.Close()

		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, err = file.Write(b)
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}

func ReadNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
		return nil
	}
	return bytes
}

func ReadUserBinary() {
	file, err := os.Create("output(2).bin")
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	defer file.Close()

	m := Website1{}
	for i := 0; i <= 10; i++ {
		data := ReadNextBytes(file, 8)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m) //二进制解码
		if err != nil {
			fmt.Println("解码失败")
			return
		}
		fmt.Println("第", i, "个", "值为：", m)
	}
}

func main() {
	WriteUserBinary()
	ReadUserBinary()
}
