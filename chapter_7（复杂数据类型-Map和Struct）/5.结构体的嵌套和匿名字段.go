package main

import "fmt"

// 结构体嵌套
type bad_people struct {
	id, age int
	name    string
	penis   math
	egg
}

type math struct {
	length, width int
}

// 匿名字段（相当于名称和结构体相同）
type egg struct {
	math
}

func main() {
	human1 := bad_people{
		id:    1,
		age:   18,
		name:  "李子康",
		penis: math{18, 8},
		egg: egg{math{
			length: 5,
			width:  5,
		}},
	}
	fmt.Println(human1)
}
