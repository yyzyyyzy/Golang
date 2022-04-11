package main

import "fmt"

type MyInt int

type person1 struct {
	name string
	age  int
}

//可以不使用对象，可以用 _ 代替

func (i *MyInt) DoSomething(a1 int, a2 int) int { // i为Python中的self，java中的this
	return a1 + a2 + int(*i)
}

func (p *person1) DoSomething(a1 int, a2 int) int {
	return a1 + a2 + p.age
}

func main() {
	var v1 MyInt = 1

	var p1 person1 = person1{
		name: "李子康",
		age:  18,
	}
	result1 := v1.DoSomething(1, 2) //v1 = i
	result2 := p1.DoSomething(2, 3) //p1 = p
	fmt.Println(result1, result2)
}
