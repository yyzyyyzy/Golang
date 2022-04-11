package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var num int = 10
	fmt.Printf("num的数据类型为：%T\n", num)
	fmt.Println("num的数据类型为:", reflect.TypeOf(num))
	fmt.Println("num所占用的字节数为:", unsafe.Sizeof(num))
}
