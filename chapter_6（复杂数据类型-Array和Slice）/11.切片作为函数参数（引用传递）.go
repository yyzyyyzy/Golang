package main

import "fmt"

func double_value1(s []int) { //引用地址传递
	for k, v := range s {
		s[k] = v * 2
	}
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1, &slice1[0], "值翻倍")
	double_value1(slice1)
	fmt.Println(slice1, &slice1[0], "地址不变")
}
