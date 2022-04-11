package main

import "fmt"

//父类
type person struct {
	id   int
	name string
	age  int
}

//子类
type GoodPeople struct {
	*person //指针类型的匿名字段
	score   int
}

//子类
type BadPeople struct {
	*person //指针类型的匿名字段
	score   int
}

func main() {
	s1 := GoodPeople{
		person: nil,
		score:  90,
	}
	s1.person = new(person)
	(*s1.person).id = 1
	(*s1.person).name = "李子康"
	(*s1.person).age = 18
	fmt.Println(s1.id, s1.name, s1.age)
}
