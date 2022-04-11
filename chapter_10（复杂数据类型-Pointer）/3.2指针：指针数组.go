package main

import (
	"fmt"
	"reflect"
)

func main() {

	a, b, c := 10, 20, 30
	x, y, z := [3]int{1, 2, 3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}

	//指针数组
	var arr [3]*int = [3]*int{&a, &b, &c}
	fmt.Println(arr, reflect.TypeOf(arr))

	//二维指针数组
	var arr1 [3]*[3]int = [3]*[3]int{&x, &y, &z}
	fmt.Println(arr1, reflect.TypeOf(arr1))

	//指针数组的赋值
	(*arr1[1])[1] = 50
	fmt.Println(y)

	//二维指针数组的遍历
	for _, v := range arr1 {
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}

	}
}
