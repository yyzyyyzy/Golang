package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4 //直接修改堆上的变量，无法传入栈帧
	v := reflect.ValueOf(x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x) //栈上存储参数地址(逃逸到堆的)，传地址可以修改堆的参数从而改变栈上参数
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p.Elem:", p.Elem().CanSet())
}
