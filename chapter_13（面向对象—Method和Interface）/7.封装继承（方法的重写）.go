package main

import "fmt"

type Dog struct { //父类
	name string
}

type Doggy struct { //子类
	Dog
	sex string
}

func (d *Dog) PrintName() {
	fmt.Printf("%v是一只小狗\n", d.name)
}

func (d *Doggy) PrintName() {
	fmt.Printf("它是一只小%v狗\n", d.sex)
}

func main() {
	d1 := Doggy{
		Dog: Dog{name: "旺财"},
		sex: "母",
	}
	d1.Dog.PrintName() //方法的重写：父类的使用场景和子类的使用场景不相同，可以在子类定义一个重名的方法改变其用法
}
