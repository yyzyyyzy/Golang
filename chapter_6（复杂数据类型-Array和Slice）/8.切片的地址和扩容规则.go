package main

import "fmt"

func slice_addr() {
	var slice1 []int
	fmt.Printf("slice1的内存地址为 %p\n", slice1)

	//append追加slice的数据时，切片地址的内存地址会发生改变
	slice1 = append(slice1, 1, 2, 3, 4)
	fmt.Println("slice的内存地址为", &slice1[0])
}

func slice_cap() {

	//slice扩容原理：
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))
}

func main() {
	slice_addr()
	slice_cap()
}
