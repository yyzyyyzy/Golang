package main

import "fmt"

func SelectMin(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func SelectSort(arr []int) []int {
	result := []int{}
	count := len(arr)
	for i := 0; i < count; i++ {
		smallest_index := SelectMin(arr)
		result = append(result, arr[smallest_index])
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}
	return result
}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(SelectSort(arr))
}
