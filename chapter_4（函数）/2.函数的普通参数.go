package main

import "fmt"

//函数定义，参数定义
func test1(a int, b string) {
	fmt.Printf("a=%d,b=%s\n", a, b)
}

//函数定义，参数定义
func test2(c, d int) {
	sum := c + d
	fmt.Println(sum)
}
func main() {
	// 函数调用，参数传递
	test1(1, "a")
	test2(3, 4)
}
