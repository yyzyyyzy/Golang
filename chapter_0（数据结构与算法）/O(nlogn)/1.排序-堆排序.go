package main

import (
	"fmt"
)

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i //每次截取一段
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}
func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := (length - 1) / 2     //二叉树的深度log2^N
	for i := depth; i >= 0; i-- { //循环所有的三结点（父节点和左右子节点）
		topmax := i                                                //假定最大的在根节点
		leftchild := 2*i + 1                                       //左子节点的索引
		rightchild := 2*i + 2                                      //右子节点的索引
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止数组越界
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}
		if topmax != i { //确保i的值为最大
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}
func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(HeapSort(arr))
}
