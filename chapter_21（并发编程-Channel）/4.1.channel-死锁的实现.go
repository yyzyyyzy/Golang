package main

import "fmt"

func main() {
	deadlock()
}

func deadlock() {
	mychan := make(chan int) //管道的数据类型，管道的缓存大小
	mychan <- 123456
	fmt.Println("写入管道的值：", <-mychan)
}
