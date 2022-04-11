package main

import "fmt"

//父类
type dog struct {
	name string
}

//子类
type doggy struct {
	dog
	sex string
}

func (d *dog) PrintInfo() {
	fmt.Printf("%v是一只小狗\n", d.name)
}

func main() {
	d1 := dog{name: "旺财"}
	d1.PrintInfo()

	d2 := doggy{dog: dog{name: "奥利奥"}, sex: "母"} //子类可以继承父类方法
	d2.PrintInfo()
}
