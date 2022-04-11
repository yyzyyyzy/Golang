package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Name string
	Age  int
}

type Boy1 struct {
	User1
	addr string
}

func SetStructValue(i interface{}) {
	i_valueof := reflect.ValueOf(i)
	i_elem := i_valueof.Elem()
	i_fieldbyname1 := i_elem.FieldByName("Name")
	i_fieldbyname2 := i_elem.FieldByName("Age")
	i_fieldbyname1.SetString("WHC")
	i_fieldbyname2.SetInt(20)
	fmt.Printf("结构体的字段被修改为%v\n", i)
}

func PrintStructFieldNum(i interface{}) { //打印结构体字段数量
	i_valueof := reflect.ValueOf(i)
	i_elem := i_valueof.Elem()
	i_numfield := i_elem.NumField()
	fmt.Printf("结构体的字段数量为%v\n", i_numfield)
}

func FindFieldByIndex(i interface{}) { //嵌套结构体需要通过fieldbyindex访问字段
	i_valueof := reflect.TypeOf(i)
	i_index := i_valueof.FieldByIndex([]int{1})
	fmt.Println(i_index.Name)

}

func main() {
	u := User1{"LZK", 18}
	PrintStructFieldNum(&u)
	SetStructValue(&u)

	b := Boy1{User1{"LZK", 18}, "南京"}
	FindFieldByIndex(b)
}
