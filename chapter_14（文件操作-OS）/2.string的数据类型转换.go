package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串——>字符切片(输出的是ASCII码)
	fmt.Println([]byte("helloworld"))

	//字符切片——>字符串
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}))

	//其他类型——>字符串
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(10, 2)) //（数据，进制）
	fmt.Println(strconv.FormatFloat(3.1415926, 'f', 4, 64))
	fmt.Println(strconv.Itoa(100))

	//字符串——>布尔类型
	bool1, err1 := strconv.ParseBool("true")
	if err1 != nil {
		fmt.Println("类型转换错误")
	} else {
		fmt.Println(bool1)
	}
	//字符串——>数字
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	int1, err2 := strconv.ParseInt("abc", 16, 64)
	if err2 != nil {
		fmt.Println("类型转换错误")
	} else {
		fmt.Println(int1)
	}

	//其他类型——>字符串，添加到字符切片
	fmt.Println(string(strconv.AppendBool(make([]byte, 0, 1024), false)))
	fmt.Println(string(strconv.AppendInt(make([]byte, 0, 1024), 10, 2)))
	fmt.Println(string(strconv.AppendFloat(make([]byte, 0, 1024), 3.14, 'f', 4, 64)))
	fmt.Println(string(strconv.AppendQuote(make([]byte, 0, 1024), "hello")))
}
