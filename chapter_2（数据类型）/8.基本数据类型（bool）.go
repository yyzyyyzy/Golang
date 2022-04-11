package main

import "fmt"

func main() {
	var a bool //零值为false，不支持数据类型转换
	fmt.Println(a)

	b := true
	fmt.Println(b)

	c := (1 == 2)
	fmt.Println(c)
}
