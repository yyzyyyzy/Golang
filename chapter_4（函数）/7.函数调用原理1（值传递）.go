package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
}

// 通过main函数调用swap函数交换a,b两个局部变量的值;
// 被调用者的参数和返回值，都在调用者的函数栈帧中;
// 传参数就是值拷贝，入栈顺序为从右往左b=2,a=1;
// 最终swap调用只改变了形式参数b=1,a=2,而main函数的局部变量没有改变
func main() {
	a, b := 1, 2
	swap(a, b)
	fmt.Println(a, b)
}
