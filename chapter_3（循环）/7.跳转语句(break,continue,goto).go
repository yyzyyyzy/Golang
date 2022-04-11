package main

import "fmt"

func main() {
	//break 跳出循环 多层嵌套则跳出最近的内循环

	//continue 跳出本次循环，进入下一次循环（只能在for循环使用）

	//goto 跳转
	fmt.Println("aaa")
	goto CCC
	fmt.Println("bbb")
CCC:
	fmt.Println("ccc")

}
