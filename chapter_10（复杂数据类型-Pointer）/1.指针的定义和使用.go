package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 指针的意义：创建一个变量值的引用
	var name string = "李子康"
	var pointer *string = &name

	fmt.Printf("变量值为：%v，地址为：%v，数据类型为%v，占用字节个数为%v\n", name, &name, reflect.TypeOf(name), unsafe.Sizeof(name))
	fmt.Printf("变量值为：%v，地址为：%v，数据类型为%v，占用字节个数为%v\n", pointer, &pointer, reflect.TypeOf(pointer), unsafe.Sizeof(pointer))

	// 空指针：nil
	var emptyPointer *int
	fmt.Println(emptyPointer)

	// 野指针：指向未知的内存地址，会报错
	//var wildPointer *int
	//*wildPointer = 88
	//fmt.Println(wildPointer)

	//指针赋值
	pointer1 := new(int)
	*pointer1 = 999
	fmt.Println(*pointer1)
}
