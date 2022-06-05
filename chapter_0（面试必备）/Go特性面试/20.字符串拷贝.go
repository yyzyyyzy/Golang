package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "12345"
	str_copy := make([]byte, len(str))
	copy(str_copy, str)
	result := string(str_copy)
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

}
