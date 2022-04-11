package main

import "fmt"

func swap1() { //方法一：两个变量进行交换，使用第三个变量
	x, y := 10, 20
	var z int
	fmt.Printf("x=%d,y=%d,z=%d\n", x, y, z)
	z = x
	x = y
	y = z
	fmt.Printf("x=%d,y=%d,z=%d\n", x, y, z)
}
func init() { //方法二：两个变量进行交换
	i := 10
	j := 20
	i, j = j, i
	fmt.Printf("i=%d,j=%d\n", i, j)
}
func main() { //错误写法：未能交换数据
	a, b := 10, 20
	a = b
	b = a
	fmt.Printf("a=%d,b=%d", a, b)
}
