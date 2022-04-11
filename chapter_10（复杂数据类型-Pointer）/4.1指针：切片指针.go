package main

import (
	"fmt"
	"reflect"
)

func main() {
	//切片指针
	slice1 := []int{1, 2, 3, 4, 5}
	slice1Pointer := &slice1
	fmt.Println(slice1Pointer, reflect.TypeOf(slice1Pointer))

	//切片指针的赋值
	(*slice1Pointer)[2] = 20
	fmt.Println(slice1Pointer)
}
