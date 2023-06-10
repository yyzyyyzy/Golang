package main

import "fmt"

// 选择排序：
// 时间复杂度：O(n^2)
// 1.在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
// 2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
// 3.重复第二步，直到所有元素均排序完毕

func SelectionSort(arr []int) []int {
	var sorted []int
	var count = len(arr)
	for i := 0; i < count; i++ {
		min := selectMin(arr)
		sorted = append(sorted, arr[min])
		arr = append(arr[:min], arr[min+1:]...)
	}
	return sorted
}

func selectMin(arr []int) int {
	min := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[min] {
			min = i
		}
	}
	return min
}

// 选择排序改进版：
// 1.每次循环找到最小值，在内循环结束后再交换

func SelectionSort_NB(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 10}
	fmt.Println(SelectionSort(arr))
	fmt.Println(SelectionSort_NB(arr))
}
