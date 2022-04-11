package main

import "fmt"

func bubblesort(slice *[]int) {
	for i := 0; i < len(*slice)-1; i++ {
		for j := 0; j < len(*slice)-i-1; j++ {
			if (*slice)[j] > (*slice)[j+1] {
				(*slice)[j], (*slice)[j+1] = (*slice)[j+1], (*slice)[j]
			}
		}
	}
}

func main() {
	slice := []int{10, 2, 30, 4, 50}
	bubblesort(&slice)
	fmt.Println(slice)

}
