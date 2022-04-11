package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 10, 10)
	copy(s2, s1)
	fmt.Println(s1, s2, &s1[0], &s2[0]) //copy()会重新开辟内存，copy()后修改不会影响到底层数组

	s1[2] = 100
	fmt.Println(s1, s2)
}
