package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
*/
func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	step1, step2 := 1, 2
	for i := 3; i <= n; i++ {
		step1, step2 = step2, step1+step2
	}
	return step2
}
