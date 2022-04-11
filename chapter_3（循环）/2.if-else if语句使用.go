package main

import "fmt"

func main() {
	if a := 10; a >= 11 {
		fmt.Println("数据溢出")
	} else if a == 10 {
		fmt.Println("数据已满")
	} else if a <= 5 {
		fmt.Println("数据过小")
	}
}
