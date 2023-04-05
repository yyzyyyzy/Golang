package main

import "fmt"

func main() {
	fmt.Println(twoSum1([]int{1, 2, 3}, 5))
	fmt.Println(twoSum2([]int{1, 5, 8, 2}, 10))

}

// 暴力
func twoSum1(nums []int, target int) []int {
	for i, v := range nums { //i第一个索引的位置
		for j := i + 1; j < len(nums); j++ { //j第二个索引的位置
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 哈希表
func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums { //i第二个数的索引
		if p, ok := hash[target-v]; ok { //p第一个数的索引
			return []int{p, i}
		}
		hash[v] = i
	}
	return nil
}
