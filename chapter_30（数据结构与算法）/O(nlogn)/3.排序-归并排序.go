package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	left_index := 0
	right_index := 0
	result := []int{}
	for left_index < len(left) && right_index < len(right) {
		if left[left_index] < right[right_index] {
			result = append(result, left[left_index])
			left_index++
		} else if left[left_index] > right[right_index] {
			result = append(result, right[right_index])
			right_index++
		} else {
			result = append(result, left[left_index])
			result = append(result, right[right_index])
			left_index++
			right_index++
		}
	}
	for left_index < len(left) {
		result = append(result, left[left_index])
		left_index++
	}
	for right_index < len(right) {
		result = append(result, right[right_index])
		right_index++
	}
	return result
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))
}
