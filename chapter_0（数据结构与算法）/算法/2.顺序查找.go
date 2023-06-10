package main

// 顺序查找
// 时间复杂度：O(n)
// 1.从第一个元素开始逐个比对，直到找到相等的值
// 2.如果没有找到相等的值，返回-1

func SequentialSearch(arr []int, key int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}
	return -1
}
