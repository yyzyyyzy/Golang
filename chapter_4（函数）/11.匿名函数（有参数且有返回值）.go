package main

import "fmt"

func main() {

	var num int = 10
	f1 := func(a int, b int) int {
		result := a + b + num
		return result
	}
	fmt.Println(f1(1, 2))
}
