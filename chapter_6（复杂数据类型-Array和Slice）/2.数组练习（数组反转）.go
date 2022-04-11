package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 2}
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
	fmt.Println("反转后的数组为 ", arr)
}
