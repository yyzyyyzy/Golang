package main

import "sort"

/*
给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

/*
暴力解法：

	遍历数组，统计每个元素出现的次数，返回出现次数大于⌊ n/2 ⌋的元素
*/
func majorityElement(nums []int) int {
	l := len(nums) / 2
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for j, k := range m {
		if k > l {
			return j
		}
	}
	return 0
}

/*
摩尔投票法：

	假设当前数为众数, 遇到相同的数字则加一,否则减一
	若前n个票和为0(互相抵消), 则设下一个数为当前数
	重复1, 2 最后的当前数为众数
*/
func majorityElement(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}

/*
排序法：

	排序后，中间的数一定是众数
*/
func majorityElement(nums []int) int {
	// 排序法
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	println(majorityElement(nums))
}
