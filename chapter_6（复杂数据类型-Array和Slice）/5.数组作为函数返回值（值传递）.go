package main

import "fmt"

func Bubblesort2(arr [10]int) [10]int { //如果想通过函数计算结果改变实参的值，那么需要将数组作为函数的返回值
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	return arr
}

func main() {
	array := [10]int{3, 98, 2, 6, 5, 4, 1, 54, 10, 9}
	array = Bubblesort2(array)
	fmt.Println(array)
}
