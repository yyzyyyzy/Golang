package main

import "fmt"

func FUCK(a, b int) (int, int, int) { //匿名返回值
	var c, d int
	sum := a - b
	c = 4
	d = 3
	return c, d, sum
}
func main() {
	var result int
	_, _, result = FUCK(11, 2)
	fmt.Println(result)
}
