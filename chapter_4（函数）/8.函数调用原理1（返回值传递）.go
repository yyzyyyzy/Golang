package main

import "fmt"

func incr1(a int) int { //匿名返回值b
	var b int

	defer func() {
		a++
		b++
	}()

	a++
	b = a
	return b
}

func main() {
	var a, b int
	b = incr1(a)
	fmt.Println(a, b)
}
