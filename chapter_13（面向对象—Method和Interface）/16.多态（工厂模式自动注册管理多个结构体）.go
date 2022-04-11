package main

import "fmt"

type Class interface {
	Do()
}

var factoryByName = make(map[string]func() Class) //保存注册好的工厂信息

func Register(name string, factory func() Class) { //注册一个类生成工厂
	factoryByName[name] = factory
}

func Create(name string) Class { //根据名称创建对应的类
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}

type Class1 struct {
}

func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {
	Register("Class1", func() Class { //启动时注册class1工厂
		return new(Class1)
	})
}

type Class2 struct {
}

func (c *Class2) Do() {
	fmt.Println("Class2")
}

func init() {
	Register("Class2", func() Class { //启动时注册class2工厂
		return new(Class2)
	})
}

func main() {
	c1 := Create("Class1")
	c2 := Create("Class2")
	c1.Do()
	c2.Do()
}
