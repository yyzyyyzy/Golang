package main

import (
	"fmt"
	"reflect"
)

type Person struct {
}

func main() {
	P1 := Person{}
	TypeofP1 := reflect.TypeOf(P1)
	fmt.Println(TypeofP1) //非指针传入，直接返回实际类型

	P2 := &Person{}
	TypeofP2 := reflect.TypeOf(P2)
	fmt.Println(TypeofP2.Elem()) //指针传入，需要使用Elem()方法得到实际值

}
