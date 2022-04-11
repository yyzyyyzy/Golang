package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a int = -10
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a)) //查看数据类型
	fmt.Println(unsafe.Sizeof(a))  //查看数据占据字节数
}
