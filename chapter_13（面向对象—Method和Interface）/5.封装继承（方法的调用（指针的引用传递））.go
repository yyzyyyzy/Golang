package main

import (
	"fmt"
)

type student struct {
	score int
}

type learner struct {
	score int
}

func (s *student) Print() {
	s.score = 100
	fmt.Println("p1方法调用结果（引用传递）：", *s)
}
func (l learner) Print() {
	l.score = 50
	fmt.Println("p2方法调用结果（值传递）：", l)
}
func main() {
	p1 := student{score: 18}
	p2 := learner{score: 19}

	p1.Print()
	fmt.Println("p1方法调用结果（引用传递）：", p1)

	p2.Print()
	fmt.Println("p2方法调用结果（值传递）：", p2)
}
