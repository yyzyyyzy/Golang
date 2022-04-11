package main

import "fmt"

func main() {

	//参数传递
	func(a, b int) { //匿名函数带参数
		var sum int
		sum = a + b
		fmt.Println(sum)
	}(3, 7) //调用时带参数

}
