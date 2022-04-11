package main

import "fmt"

func Swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
}

func main() {
	num1 := 10
	num2 := 20
	Swap(&num1, &num2)
	if num1 == 20 && num2 == 10 {
		fmt.Println("指针作为函数参数传递是引用传递，形参可以改变实参")
	}
}
