package main

import "fmt"

// 结构体嵌套
// 父类
type Person struct {
	id   int
	name string
	age  int
}

// 子类
type Student struct {
	Person //匿名字段  只有类型  没有名字
	score  int
}

func main() {
	//实例1
	student1 := Student{
		Person: Person{
			id:   1,
			name: "李子康",
			age:  18,
		},
		score: 100,
	}
	fmt.Println(student1)

	//实例2
	student2 := Student{
		Person: Person{
			id:   0,
			name: "",
			age:  0,
		},
		score: 90,
	}
	fmt.Println(student2)

}
