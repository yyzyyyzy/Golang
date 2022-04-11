package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("E:\\函")
	if err != nil {
		fmt.Println("出现问题", err)
	}
	fmt.Println(file, err)
	fmt.Println(reflect.TypeOf(err))
}
