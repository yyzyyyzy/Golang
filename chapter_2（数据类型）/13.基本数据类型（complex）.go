package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := 3.2 + 3.14i
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(reflect.TypeOf(a))

	b := 1 + 2i
	fmt.Println(real(b), imag(b)) //取实部real(),取虚部imag()
}
