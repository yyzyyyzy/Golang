package main

import "fmt"

type Pig interface { //子集
	Piggo()
	Pigspeak()
}

type Piggy interface { //超集:超集的方法子集不能使用
	Pig //匿名字段实现继承
	Pigeat()
}

type horse struct {
}

func (h *horse) Piggo() {
	fmt.Println("迈着猪八戒的步伐")
}

func (h *horse) Pigspeak() {
	fmt.Println("哼哧哼哧叫")
}

func (h *horse) Pigeat() {
	fmt.Println("呼噜呼噜地吃")
}

func Pigcando(p1 Pig, p2 Piggy) { //接口使用的关键
	p1.Pigspeak()
	p1.Piggo()
	p2.Pigeat()
}

func main() {
	h := horse{}
	Pigcando(&h, &h)
}
