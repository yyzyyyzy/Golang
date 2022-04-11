package main

import (
	"fmt"
)

var mymap = make(map[int]int)

func factorialFor(num int) (ret int) {
	ret = 1
	for i := 1; i <= num; i++ {
		ret *= i
	}
	mymap[num] = ret
	return ret
}

func main() {
	for i := 1; i <= 10; i++ {
		factorialFor(i)
		//go factorialFor(i) //map不是线程安全的，需要管道处理线程安全问题
	}
	//time.Sleep(time.Second)
	for k, v := range mymap {
		fmt.Printf("%d的阶乘为：%d\n", k, v)
	}
}
