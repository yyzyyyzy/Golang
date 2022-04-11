package main

import "fmt"

func testA() {
	fmt.Println("testA")
}

func testB(i int) {
	//需要在defer函数中使用recover函数！！
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [3]int
	arr[i] = 999
}

func testC() {
	fmt.Println("testC")
}

func main() {
	testA()
	testB(3) //触发异常
	testC()
}
