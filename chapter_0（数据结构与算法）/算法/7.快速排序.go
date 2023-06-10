package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := arr[0]
	var left, right []int
	for _, value := range arr[1:] {
		if value <= mid {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return append(QuickSort(left), append([]int{mid}, QuickSort(right)...)...)
}

// QuickSortV2 改进版(三数取中值)
func QuickSortV2(arr []int, low, hight int) {
	if low >= hight {
		return
	}

	left, right := low, hight
	pivot := arr[(low+hight)/2] // 这里的经验值取的是中间数，经过 Benchmark 测试，确实比较优秀

	for left <= right {
		// 从左边开始迭代

		// 左边的数如果比 pivot 小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for arr[left] < pivot {
			left++
		}

		// 右边的数如果比 pivot 大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for arr[right] > pivot {
			right--
		}

		// 这里进行一次交换，将上面碰到的大数和小数交换一次
		//left 继续右走，right 继续左走 注意这里还不一定相遇，去继续执行上面的逻辑
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	// 【 xxx[xxxxx]xxxxxx】
	// 【 xxxxxx][xxxxxxxx】
	// [] => left,right
	// 【】 => low,hight
	QuickSortV2(arr, low, right)
	QuickSortV2(arr, left, hight)
}

// QuickSortV3 改进版
func QuickSortV3(arr []int, left, right int) {
	if left >= right {
		return
	}
	cur, lo := left+1, left
	for cur <= right {
		if arr[cur] <= arr[left] {
			arr[lo+1], arr[cur] = arr[cur], arr[lo+1]
			lo++
		}
		cur++
	}
	arr[left], arr[lo] = arr[lo], arr[left]
	QuickSortV3(arr, left, lo-1)
	QuickSortV3(arr, lo+1, right)
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(QuickSort(arr))
}
