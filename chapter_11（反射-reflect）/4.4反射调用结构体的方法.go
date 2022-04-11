package main

import (
	"fmt"
	"reflect"
)

type User3 struct {
	Id   int
	Name string
	Age  int
}

func (u User3) FuncHasArgs(id int, name string, age int) {
	fmt.Println("反射调用有参数的方法 id:", id, ",name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User3) FuncNoArgs() {
	fmt.Println("反射调用无参数的方法")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
	user := User3{0, "Wang.H.C", 25}

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	valueof := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := valueof.MethodByName("FuncHasArgs")
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("Li.Z.K"), reflect.ValueOf(20)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = valueof.MethodByName("FuncNoArgs")
	args = []reflect.Value{}
	methodValue.Call(args)
}
