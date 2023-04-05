package main

import (
	"fmt"
	"reflect"
)

// 声明一下结构体
type Test struct {
	Name string
	Age  int8
}

func (this Test) GetName(name string) {
	fmt.Println(name)
}

func (this *Test) GetName1(name string) {
	fmt.Println(name)
}

func main() {

	//实例化一个结构体
	str := Test{Name: "will", Age: 12}

	//第一个传值操作，这里直接传个结构体
	//获取变量本身类型的信息
	testType := reflect.TypeOf(str)
	//获取变量类型名称
	testTypeName := testType.Name()
	fmt.Println("typeName: " + testTypeName)
	//获取结构的字段数量
	fmt.Println("numFiled :", testType.NumField())
	//获取结构的方法数量
	fmt.Println("numMethod :", testType.NumMethod())

	//获取变量实例本身类型的信息
	v := reflect.ValueOf(str)
	//调用方法，并传参数
	param := []reflect.Value{reflect.ValueOf("will")}
	v.MethodByName("GetName").Call(param)
	//获取结构的字段数量
	fmt.Println("NumFiled: ", v.NumField())
	//获取变量类型名称
	fmt.Println("NumFiled: ", v.Type().Name())

	//第二个传指针操作，这里直接传个结构体指针
	//注意这里，如果是指针的话，要使用Elem方法来获取对象
	testType1 := reflect.TypeOf(&str).Elem()
	testTypeName1 := testType.Name()
	fmt.Println("typeName: " + testTypeName1)
	fmt.Println("numFiled :", testType1.NumField())
	fmt.Println("numMethod :", testType1.NumMethod())

	//注意这里，如果是指针的话，要使用Elem方法来获取对象
	v1 := reflect.ValueOf(&str).Elem()
	param1 := []reflect.Value{reflect.ValueOf("willd")}
	v1.MethodByName("GetName").Call(param1)
	fmt.Println(str)

	//设置字段的值，首先判断一下，值的类型，是否可以设置
	if v1.Kind() == reflect.Ptr && !v1.CanSet() {
		fmt.Println("can not set")
		return
	}
	v1.FieldByName("Name").SetString("test2")
	v1.FieldByName("Age").SetInt(23)

	fmt.Println(str)
	//看完以上的实例，这里要注意一下，如果是想修改字段的值，这里必须传指针，因为go语言中，基本的类型，在传值时，都是按值传递的，也是就我们说的会复制一份。

}
