package main

import (
	"fmt"
	"sort"
)

type MyintList []int

func (m MyintList) Len() int {
	return len(m)
}

func (m MyintList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyintList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	myintlist := MyintList{9, 11, 3, 10, 5, 6, 88, 8, 1}

	sort.Sort(myintlist) //sort包内的sort方法是一个封装好的快排
	fmt.Println(myintlist)
}
