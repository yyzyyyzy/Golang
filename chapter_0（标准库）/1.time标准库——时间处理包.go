package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	fmt.Println(time.Now().Add(time.Hour))

	//不能用int类型和time.Duration类型做运算
	num := 1
	time.Sleep(time.Duration(num) * time.Second)
}
