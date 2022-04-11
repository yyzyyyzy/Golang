package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 10

	// 一级指针
	numPtr := &num
	fmt.Println(numPtr, reflect.TypeOf(numPtr))

	// 二级指针
	numPPtr := &numPtr
	fmt.Println(numPPtr, reflect.TypeOf(numPPtr))

	// 三级指针
	numPPPtr := &numPPtr
	fmt.Println(numPPPtr, reflect.TypeOf(numPPPtr))

	//多级指针的互相修改
	num1 := 20
	*numPPtr = &num1
	fmt.Println(*numPtr)
}
