package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num) //接口类型变量 --> reflection.Type/reflection.Value
	value := reflect.ValueOf(num)    //reflect.Value类型对象

	fmt.Println(reflect.TypeOf(pointer))
	fmt.Println(reflect.TypeOf(value))

}
