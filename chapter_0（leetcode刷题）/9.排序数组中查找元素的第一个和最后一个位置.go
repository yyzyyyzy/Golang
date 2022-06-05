package main

import (
	"sort"
)

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target) //SearchInts()方法会查找到切片指定的索引,没找到就返回元素个数
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
