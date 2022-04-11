package main

import "fmt"

func test1(i int) {
	var arr [3]int
	arr[i] = 999
	fmt.Println(arr)
}
func main() {
	test1(4) //数组越界：panic: out of range
}
