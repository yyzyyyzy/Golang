package main

import "fmt"

func main() {
	//数组指针
	arr := [3]int{1, 2, 3}
	arrPointer := &arr
	fmt.Println(arrPointer, *arrPointer)

	//数组指针的赋值
	(*arrPointer)[0] = 111
	fmt.Println(arrPointer, *arrPointer)

	//数组指针的遍历
	for k, v := range arrPointer {
		fmt.Println(k, v)
	}
}
