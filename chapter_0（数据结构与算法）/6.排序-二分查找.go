package main

import "fmt"

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(arr)
	fmt.Println(quickSort(arr))
	fmt.Println(BinarySearch(quickSort(arr), 5))
}

func BinarySearch(arr []int, target int) int {
	low := 0             //最小索引
	high := len(arr) - 1 //最大索引
	for low < high {
		mid := (low + high) / 2
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := arr[0]
	var left, right []int
	for _, index := range arr[1:] {
		if index <= mid {
			left = append(left, index)
		} else {
			right = append(right, index)
		}
	}
	return append(quickSort(left), append([]int{mid}, quickSort(right)...)...)
}
