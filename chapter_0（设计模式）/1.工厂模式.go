package main

import "fmt"

type Restaurant interface {
	GetFood()
}

type Goubuli struct {
}

func (g *Goubuli) GetFood() {
	fmt.Println("狗不理包子生产完毕，继续。。。")
}

type Beijingkaoya struct {
}

func (b Beijingkaoya) GetFood() {
	fmt.Println("北京烤鸭生产完毕，继续。。。")
}

// 生产工厂
func NewRestaurant(name string) Restaurant {
	switch name {
	case "goubuli":
		return &Goubuli{}
	case "Beijingkaoya":
		return &Beijingkaoya{}
	}
	return nil
}

func main() {
	NewRestaurant("goubuli").GetFood()
	NewRestaurant("Beijingkaoya").GetFood()
}
