package main

import (
	"fmt"
)

//获取一个数组里最大值，并且拿到下标
//获取一个数组里最小值，并且拿到下标
func main() {

	//声明一个数组5个元素
	arr := [5]int{1, 2, 3, 4, 5}
	//假设第一个元素是最大值，下标为0
	maxVal := arr[0]
	maxIndex := 0

	for i := 1; i < len(arr); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < arr[i] {
			maxVal = arr[i]
			maxIndex = i
		}
	}

	fmt.Printf("最大值为%v, 其下标为%v\n", maxVal, maxIndex)

	arr2 := [5]int{10, 20, 30, 40, 50}
	minVal := arr2[0]
	minIndex := 0

	for i := 1; i < len(arr); i++ {
		if minVal > arr2[i] {
			minVal = arr2[i]
			minIndex = i
		}
	}
	fmt.Printf("最小值为%v, 其下标为%v\n", minVal, minIndex)

	sum := maxVal + minVal
	fmt.Printf("最大值与最小值之和为%v", sum)
}
