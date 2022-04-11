package main

import "fmt"

func demo(s []int) []int {
	s = append(s, 6, 7, 8, 9, 10) //当原切片的容量不够存储append新增加的元素，那么会在堆区新开辟内存空间并拷贝元素
	fmt.Println(s)
	return s
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = demo(slice) //demo函数的参数栈帧指向堆区新的内存地址
	fmt.Println(slice)
}
