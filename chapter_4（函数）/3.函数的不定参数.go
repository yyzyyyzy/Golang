package main

import "fmt"

//求0到5的累加
func sum(args ...int) { //固定类型的不定参数
	sum := 0
	for _, v := range args {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	sum(1, 2, 3, 4, 5)
}
