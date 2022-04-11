//类型转换
//低类型向高类型转换（float32-->float64）
package main

import "fmt"

func main() {

	a, b, c := 10, 23, 34
	sum := a + b + c
	fmt.Println(float64(sum) / 3) //int转换成float类型，只可以转换兼容类型（数字和布尔类型没办法转化）
}
