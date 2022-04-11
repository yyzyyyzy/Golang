package main

import "fmt"

type MultiN struct {
	num    int
	result int
}

func (m *MultiN) FactorialFor1() {
	ret := 1
	for i := 1; i <= m.num; i++ {
		ret *= i
	}
	m.result = ret

}
func main() {
	mychan := make(chan *MultiN)

	for i := 1; i <= 10; i++ {
		go func(n int, mychannel chan *MultiN) {
			m := &MultiN{n, 1}
			m.FactorialFor1()
			mychannel <- m
		}(i, mychan)
	}

	for i := 1; i <= 10; i++ {
		data := <-mychan
		fmt.Println(data)
	}

}
