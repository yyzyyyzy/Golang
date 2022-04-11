package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string `json:"name" NBA:"干就完了"`
	Age  int    `json:"age"`
}

func GetStructTag(i interface{}) {
	i_typeof := reflect.TypeOf(i)
	i_elem := i_typeof.Elem()

	for i := 0; i < i_elem.NumField(); i++ {
		i_Tag := i_elem.Field(i).Tag.Get("json")
		fmt.Printf("%s\n", i_Tag)
	}

}

func main() {
	t := T{}
	GetStructTag(&t)
}
