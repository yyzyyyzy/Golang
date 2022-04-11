package main

import (
	"fmt"
)

type Product interface {
	create()
}

type Product1 struct {
}

func (p1 Product1) create() {
	fmt.Println("这是产品1")
}

type Product2 struct {
}

func (p2 Product2) create() {
	fmt.Println("这是产品2")
}

type Factory struct {
}

func (f Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		return nil
	}
}

func main() {
	//创建一个工厂类
	factory := Factory{}

	p1 := factory.Generate("product1")
	p1.create()

	p2 := factory.Generate("product2")
	p2.create()
}
