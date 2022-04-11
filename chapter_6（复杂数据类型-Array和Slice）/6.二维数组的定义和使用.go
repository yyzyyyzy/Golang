package main

import "fmt"

func main() {
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}} //[]为二维数组的行[]为二维数组的列
	fmt.Println(array)

	//行数
	fmt.Println(len(array)) //此时输出的不是元素个数，而是数组的行数

	//列数
	fmt.Println(len(array[0])) //此时输出的不是元素，而是数组的列数

	//赋值
	array[0][1] = 123
	array[1][2] = 234
	fmt.Println(array)

	//二维数组的取值（for循环嵌套）
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			fmt.Println(array[i][j])
		}
	}
}
