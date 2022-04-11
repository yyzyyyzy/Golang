package main

import (
	"fmt"
	"reflect"
)

func Add(v int) int {
	return v + 10
}

func main() {

	valueof := reflect.ValueOf(Add)

	args := []reflect.Value{reflect.ValueOf(Add(2))}

	fmt.Println(valueof.Call(args)[0].Int())
}
