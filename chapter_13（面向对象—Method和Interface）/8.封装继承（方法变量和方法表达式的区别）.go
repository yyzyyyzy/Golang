package main

import (
	"fmt"
)

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}

func GetFunc() func() string {
	a := A{name: "lzk in GetFunc"}
	return a.GetName
}

func main() {
	a := A{name: "lzk in main"}
	f2 := a.GetName
	fmt.Println(f2()) //这是方法表达式，显示调用，输出lzk in main

	f3 := GetFunc()
	fmt.Println(f3()) //这是方法变量，捕获了接收者a的FunctionValue，输出lzk in GetFunc

}
