package main

import "fmt"

func init() {
	a := 10
	b := 10
	// IDE自动推导数据类型方法 num := 10
	fmt.Printf("a=%d,b=%d", a, b)
}

func main() {
	var num int = 10
	fmt.Println("num=", num)
}
