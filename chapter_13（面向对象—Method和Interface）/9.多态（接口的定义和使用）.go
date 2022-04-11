package main

import "fmt"

type CatBehavior interface {
	Catsay()
	Catgo()
}

type Dog1 struct {
}

func (d Dog1) Catsay() {
	fmt.Println("喵喵喵")
}

func (d Dog1) Catgo() {
	fmt.Println("走猫步")
}

func Catcando(C CatBehavior) { //接口使用的关键
	C.Catsay()
	C.Catgo()
}

func main() {
	d := Dog1{}
	Catcando(d)
}
