package main

import (
	"fmt"
)

func slice_define() {

	//切片包含三个部分，元素存哪里（默认零值nil），存了多少个元素（长度len），可以存多少个元素（容量cap），本质上就是数组
	var slice1 []int
	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Println(slice1, len(slice1), cap(slice1))

	//make函数定义切片（make(Type, len, cap)）
	slice2 := make([]int, 5, 10)
	slice2 = append(slice2, 1)
	slice2[0] = 1
	fmt.Println(slice2, len(slice2), cap(slice2))

	//slice的循环赋值
	slice3 := make([]int, 5, 10)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = i + 1
	}
	fmt.Println(slice3, len(slice3), cap(slice3))

}

func main() {
	slice_define()
}
