package main

import "fmt"

type hero struct {
	name  string
	age   int
	power int
}

func define_struct_map_slice() {

	//将结构体作为map的key或者value
	m1 := make(map[int]hero)

	m1[110] = hero{
		name:  "钢铁侠",
		age:   35,
		power: 100,
	}

	m1[120] = hero{
		name:  "美国队长",
		age:   25,
		power: 80,
	}
	fmt.Println(m1)

	//将切片结构体作为map的key和value
	m2 := make(map[int][]hero)

	m2[110] = []hero{{name: "绿灯侠", age: 45, power: 55}, {name: "猫女", age: 23, power: 20}}
	m2[110] = append(m2[110], hero{name: "冬兵", age: 30, power: 75})
	fmt.Println(m2)

}

func main() {
	define_struct_map_slice()
}
