package main

import (
	"fmt"
	"reflect"
)

// 空接口可以接收任意类型的数据（int, float, string, array, slice, map, struct, pointer.....）
func main() {
	var i interface{}
	fmt.Println(reflect.TypeOf(i))

	i = map[int]string{1: "LZK", 2: "WHC"}
	fmt.Println(reflect.TypeOf(i))

	i = 3.14
	fmt.Println(reflect.TypeOf(i))

}
