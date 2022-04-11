package main

import "fmt"

func caonima2(arr ...int) {
	fmt.Println(arr)
}
func caonima1(arr ...int) {
	//传递指定的数据(使用索引)
	caonima2(arr[2:5]...)
	//传递所有的数据
	caonima2(arr...)
}
func main() {
	caonima1(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
