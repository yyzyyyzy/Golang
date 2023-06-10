package main

import "fmt"

// 冒泡排序：
// 时间复杂度：O(n^2)
// 1.比较相邻的元素，如果第一个比第二个大，就交换它们两个
// 2.对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对，这样最后的元素应该会是最大的数
// 3.针对所有的元素重复以上的步骤，除了最后一个

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 冒泡排序改进版：
// 插入标志位：如果一次内循环中没有发生交换，说明已经排好序了，就不需要再循环了

func Bubblesort_NB(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		swap := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 10}
	//fmt.Println(BubbleSort(arr))
	fmt.Println(Bubblesort_NB(arr))
}
