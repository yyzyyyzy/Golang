package main

//统计一个数字在排序数组中出现的次数
//输入: nums = [5,7,7,8,8,10], target = 8
//输出: 2

func search(nums []int, target int) int {
	count := 0
	for i, _ := range nums {
		if nums[i] == target {
			count++
		}
	}
	return count
}
