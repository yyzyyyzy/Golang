package main

import "fmt"

type Hero struct {
	name  string
	age   int
	power int
}

func test2(hero1 Hero) {
	hero1.name = "猫女"
	fmt.Println(hero1)
}

func main() {
	hero1 := Hero{
		name:  "钢铁侠",
		age:   35,
		power: 100,
	}
	test2(hero1)
	fmt.Println(hero1)
}
