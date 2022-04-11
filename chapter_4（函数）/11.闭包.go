package main

import "fmt"

func A(i int) {
	i++
	fmt.Println(i)
}
func B() {
	f1 := A
	f1(1)
}

func main() {
	f2 := A
	f2(1)
}
