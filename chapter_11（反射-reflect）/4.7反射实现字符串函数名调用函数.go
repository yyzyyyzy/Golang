package main

import (
	"fmt"
	"reflect"
)

func main() {
	animal := Animal{}
	value := reflect.ValueOf(animal)
	f := value.MethodByName("Fuck")
	f.Call([]reflect.Value{})
}

type Animal struct{}

func (a Animal) Eat() {
	fmt.Println("Eat")
}

func (a Animal) Fuck() {
	fmt.Println("Fuck")
}
