package main

import "fmt"

var a int = 10 //全局变量：可以与局部变量名重名

func test9() {
	a := 10
	a++
}

func test10() {
	fmt.Println(a, &a)
}

func main() {
	a := 9
	test9()
	fmt.Println(a, &a) //全局变量和局部变量重名，优先输出局部变量
	test10()
}
