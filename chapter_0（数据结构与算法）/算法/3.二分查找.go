package main

import (
	"fmt"
	"sort"
)

// 二分查找:
// 时间复杂度：O(logn)
// 1.从有序数组的中间元素开始，如果中间元素正好是要查找的元素，则搜素过程结束；
// 2.如果某一特定元素大于搜索元素，则在数组左半部分查找，如果小于则在右半部分查找；
// 3.重复第一步。
// 4.如果直到最后也没有找到，则返回-1。

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

type mySlice []int

func (s mySlice) Len() int {
	return len(s)
}

func (s mySlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s mySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var arr mySlice = []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	sort.Sort(arr)

	fmt.Println(arr)
	fmt.Println(BinarySearch(arr, 5))
}
