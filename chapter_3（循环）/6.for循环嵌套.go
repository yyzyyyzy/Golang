package main

import (
	"fmt"
)

func init() {
	//一只公鸡5元，一只母鸡3元，3只小鸡1元，想用100元买100只鸡，有多少种排列组合？
	count := 0 //循环计数器
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			chicken := 100 - cock - hen
			if chicken%3 == 0 && cock*5+hen*3+chicken/3 == 100 {
				fmt.Println(cock, hen, chicken)
			}
		}
	}
	fmt.Println(count)
}
func main() {
	// [0,1,2,3,4] [0,1,2,3,4]两两组合有多少种排列？
	count := 0 //循环计数器
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j) //两组数据随机排序的结果
			count++           //循环计数器
		}
	}
	fmt.Println(count)
}
