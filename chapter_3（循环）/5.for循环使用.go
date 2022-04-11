package main

import "fmt"

func init() {
	str := "abc"
	for j := 0; j < len(str); j++ {
		fmt.Printf("str[%d]=%c\n", j, str[j])
	}
}
func main() {
	var sum int
	for i := 1; i < 11; i++ { //计算1+2+...+10的和
		sum = sum + i
	}
	fmt.Println(sum)
}
