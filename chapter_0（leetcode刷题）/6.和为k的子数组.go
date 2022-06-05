package main

import "fmt"

//给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
//输入：nums = [1,1,1], k = 2
//输出：2

//暴力枚举
func subarraySum1(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ { //0，1，2
		sum := 0
		for j := i; j >= 0; j-- { //0，1，2
			sum += nums[j] //num[0], num[0]+num[1], num[0]+num[1]+num[2]
			if sum == k {
				count++
			}
		}
	}
	return count
}

//前缀和+递归
//pre[i]=num[0]+ ... + num[i]
//	--> pre[i] = pre[i-1] + num[i]
//		--> pre[i] - pre[j] = num[j+1] + num[j+2] + ... + num[i]
//			--> 子数组和为k 可以转化为 pre[i] - pre[j-1] = k
//				--> pre[j-1] == pre[i] - k
//					--> 所以只要统计有多少个和为 pre[i] - k 的 pre[j-1]即可
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok { //两数之和的hash思路，（前缀和减去k）如果存在hash表内
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum1([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum2([]int{1, 3, 1}, 4))
}
