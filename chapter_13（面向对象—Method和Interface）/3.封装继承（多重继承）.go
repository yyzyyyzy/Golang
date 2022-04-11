package main

import "fmt"

type human struct {
	id   int
	name string
	age  int
}

type mankind struct {
	height int
	weight int
	length int
}

type littleboy struct {
	//继承了多个父类
	human
	mankind
}

func main() {
	b1 := littleboy{
		human: human{
			id:   1,
			name: "李子康",
			age:  18,
		},
		mankind: mankind{
			height: 189,
			weight: 150,
			length: 18,
		},
	}
	fmt.Println(b1)
}
