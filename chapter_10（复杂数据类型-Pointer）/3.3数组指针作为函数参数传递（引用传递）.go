package main

import "fmt"

func bubble_sort(arrpointer *[5]int) {
	for i := 0; i < len(*arrpointer)-1; i++ {
		for j := 0; j < len(*arrpointer)-i-1; j++ {
			if (*arrpointer)[j] > (*arrpointer)[j+1] {
				(*arrpointer)[j], (*arrpointer)[j+1] = (*arrpointer)[j+1], (*arrpointer)[j]
			}
		}
	}
}

func main() {
	arr := [5]int{10, 2, 30, 4, 50}
	bubble_sort(&arr)
	fmt.Println(arr)
}
