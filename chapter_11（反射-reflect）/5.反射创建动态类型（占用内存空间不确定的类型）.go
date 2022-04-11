package main

import (
	"fmt"
	"reflect"
)

func CreateVar_float64() {
	var num float64 = 1.64

	pointer := reflect.TypeOf(num)
	New_pointer := reflect.New(pointer)

	fmt.Println(New_pointer.Type(), New_pointer.Kind())

}

func CreateVar_Struct() {
	type T struct {
		Name string
		Age  int
	}
	t := T{"Li.Z.K", 18}

	pointer := reflect.TypeOf(t)
	New_pointer := reflect.New(pointer)

	fmt.Println(New_pointer.Type(), New_pointer.Kind())
}

func main() {
	CreateVar_float64()
	CreateVar_Struct()
}
