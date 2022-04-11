package main

import "fmt"

// 冒泡排序的思路：
// 比较两个相邻元素, 通过循环确定元素的最大值和最小值
func main() {
	arr := [...]int{3, 98, 2, 6, 5, 4, 1, 54, 10}
	for i := 0; i < len(arr)-1; i++ { // 确定数组的最大值的比较次数
		for j := 0; j < len(arr)-1-i; j++ { // 数组内元素两两比较的次数
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
