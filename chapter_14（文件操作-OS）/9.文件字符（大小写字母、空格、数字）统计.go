package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type CharCount struct {
	ChCount      int
	NumCount     int
	SpaceCount   int
	ChineseCount int
	OthersCount  int
}

func CountFileBytes(srcFileName string) {
	file, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	defer file.Close()

	var Count CharCount //实例化

	bufioRead := bufio.NewReader(file)

	for {
		str, err := bufioRead.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		for _, v := range str {
			switch {
			case (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z'):
				Count.ChCount++
			case v == ' ' || v == '\t' || v == '\r' || v == '\n':
				Count.SpaceCount++
			case v >= '0' && v <= '9':
				Count.NumCount++
			case unicode.Is(unicode.Han, v):
				Count.ChineseCount++
			default:
				Count.OthersCount++
			}
		}
	}
	fmt.Printf("字符数量= %v  数字的个数 = %v 空格的数量是= %v 汉字的数量是= %v 另外其它字符的个数是= %v",
		Count.ChCount, Count.NumCount, Count.SpaceCount, Count.ChineseCount, Count.OthersCount)
}

func main() {
	path := "E:\\golandlearning\\chapter_0（底层原理）\\5.map底层数据结构？hash算法？map扩容？"
	CountFileBytes(path)
}
