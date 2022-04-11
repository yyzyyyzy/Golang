package main

import "fmt"

func Fuck(a, b int) (c, d, sum int) { //命名返回值
	sum = a + b
	c = 1
	d = 2
	return
}

func main() {
	var result int
	_, _, result = Fuck(1, 23) //此时Fuck()就是y(x)的值
	fmt.Println(result)
}
