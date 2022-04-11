package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func demo1(a, b int) {
	fmt.Println(a + b)
}

type FUNCDEMO func(int, int)

func main() {
	fmt.Println(demo1) //函数名对应一个内存地址，存在虚拟内存空间的代码区
	fmt.Println(reflect.TypeOf(demo1))
	fmt.Println(unsafe.Sizeof(demo1))
	var f FUNCDEMO
	f = demo1
	fmt.Println(f)
	f(10, 20)

}
