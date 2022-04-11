package main

import (
	"fmt"
	"reflect"
)

type S struct {
}

func main() {

	var a *int
	fmt.Println(reflect.ValueOf(a).IsZero()) //判断反射变量是否为零值
	fmt.Println(reflect.ValueOf(a).IsNil())  //判断反射指针变量是否为零值

	fmt.Println(reflect.ValueOf(S{}).FieldByName("Name").IsValid()) //IsValid判断值有效性（字段/方法是否存在）
	fmt.Println(reflect.ValueOf(S{}).MethodByName("Tell").IsValid())
}
