package main

import (
	"fmt"
	"reflect"
)

func SetValue(i interface{}) {
	i_valueof := reflect.ValueOf(i)
	i_kind := i_valueof.Kind()
	switch i_kind {
	case reflect.Float64:
		i_valueof.SetFloat(3.1415926)
		fmt.Println("Num is", i_valueof.Float())
	case reflect.Ptr:
		i_valueof.Elem().SetFloat(3.1415926)
		fmt.Printf("Num is %v addr is %v", i_valueof.Elem().Float(), i_valueof.Pointer())
	}

}

func main() {
	Num := 3.14
	SetValue(&Num)
}
