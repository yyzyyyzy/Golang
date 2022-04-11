package main

import "fmt"

func main() {
	var a string //零值为""
	fmt.Println(a)

	str1 := "你真的帅" //汉字占3个字节（字节=byte）
	str2 := "abcd"
	str3 := str1 + str2               //字符串拼接
	fmt.Println(len(str1), len(str2)) //len()计算字节个数
	fmt.Println(str3)
	fmt.Printf("%c,%c\n", str2[0], str2[1]) //通过索引操作字符串
}
