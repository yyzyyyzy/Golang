package main

import "fmt"

func incr2(a int) (b int) { //命名返回值
	defer func() {
		a++
		b++
	}()

	a++
	return a
}

func main() {
	var a, b int
	b = incr2(a)
	fmt.Println(a, b)
}
