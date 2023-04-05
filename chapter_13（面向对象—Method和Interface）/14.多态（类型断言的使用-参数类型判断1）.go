// 类型断言：由于golang接口值的类型是动态不固定的，需要看它的动态值的类型才能确定，所以需要类型断言
//
//	switch-case-default和类型断言相结合
package main

import (
	"fmt"
)

type SliceElement []interface{}

func main() {
	S := append(SliceElement{}, 10, 3.14, "LZK")
	for _, v := range S {
		switch value := v.(type) {
		case int:
			fmt.Println("int", value)
		case float64:
			fmt.Println("float64", value)
		case string:
			fmt.Println("string", value)
		default:
			fmt.Println("unknown", value)
		}
	}
}
