package main

import "fmt"

func main() {

	var num int = 10
	func() { //匿名函数
		num++
	}()
	fmt.Println(num)

}
