package main

import "fmt"

func Bubblesort1(arr [10]int) { //数组作为参数是值传递
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	array := [10]int{3, 98, 2, 6, 5, 4, 1, 54, 10, 9}
	Bubblesort1(array)
	fmt.Println(array) //由于被调用函数的内存销毁，所以此处打印的还是原来的顺序

}
