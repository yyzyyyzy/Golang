package main

import "fmt"

// 局部变量：函数内定义的变量无法在该函数外使用; 如果局部变量和全局变量重名,优先访问局部变量
func test5() {
	a := 10
	fmt.Println(a)
}

func main() {
	a := 11
	test5()        // 输出test5()的局部变量的值10，函数调用完就销毁
	fmt.Println(a) //输出main()的局部变量的值11
}
