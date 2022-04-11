package main

import (
	"fmt"
	"os"
)

//scanln是fscanln的封装
func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Fscanln(os.Stdin, &num)
}
