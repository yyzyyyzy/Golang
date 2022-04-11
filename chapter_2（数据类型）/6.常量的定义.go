package main

import "fmt"

func main() {
	//const常量：程序运行期间，不可改变的变量
	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Println(i, j)
}
