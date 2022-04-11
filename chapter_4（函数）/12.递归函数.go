package main

import "fmt"

//递归函数格式
//Go语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中
//销毁栈内的调用空间是从内向外的（后进先出），所以打印顺序也是倒置的

func test01(a int) {
	if a == 1 {
		fmt.Println(a)
		return
	}
	test01(a - 1)
	fmt.Println(a)
}

func test02(num int) int {
	if num == 100 {
		return 100
	}
	return num + test02(num+1) //实现1+...+100的累加和
}

func main() {
	test01(3)
	fmt.Println(test02(1))
}
