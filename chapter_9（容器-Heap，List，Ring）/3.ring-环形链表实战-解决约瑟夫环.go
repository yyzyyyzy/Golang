package main

import (
	"container/ring"
	"fmt"
)

type Player struct {
	pos   int
	alive bool
}

const (
	startIndex = 1
	playCount  = 100
	step       = 3
)

func main() {
	ring := ring.New(100)
	//约瑟夫环
	for i := startIndex; i <= playCount; i++ {
		ring.Value = &Player{i, true}
		ring = ring.Next()
	}
	//ring.Do(func(i interface{}) {
	//	fmt.Print(i,"\t")
	//})
	deadCount := 0
	counter := 1

	for deadCount < playCount {
		ring = ring.Next()
		if ring.Value.(*Player).alive {
			counter++
		}
		if counter == step {
			ring.Value.(*Player).alive = false
			fmt.Printf("%d号被扔进大海\n", ring.Value.(*Player).pos)
			deadCount++
			counter = 0
		}
	}

}
