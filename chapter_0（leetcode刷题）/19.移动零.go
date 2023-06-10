package main

/*

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]

*/

/*
暴力替换：
 1. 遍历数组，统计0的个数
 2. 遍历数组，将非零元素前移，后面补0
*/
func moveZeroes(nums []int) {
	index, count := 0, 0
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		} else {
			count++
		}
	}
	for i := 0; i < count; i++ {
		nums[index] = 0
		index++
	}
}

/*
快慢指针:
 1. 慢指针指向第一个零，快指针指向未知元素
 2. 如果快指针是零，直接++，否则与慢指针交换元素，然后两个指针++
*/
func moveZeroes(nums []int) {
	n := len(nums)
	for i, j := 0, 0; j < n; j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
