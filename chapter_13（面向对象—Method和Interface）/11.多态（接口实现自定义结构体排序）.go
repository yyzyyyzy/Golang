package main

import (
	"fmt"
	"sort"
)

type Teacher struct {
	Name  string
	Score int
}
type Teachers []Teacher

func (t Teachers) Len() int {
	return len(t)
}

func (t Teachers) Less(i, j int) bool {
	if t[i].Score < t[j].Score {
		return true
	} else if t[i].Score > t[j].Score {
		return false
	} else {
		return t[i].Name < t[j].Name
	}
}

func (t Teachers) Swap(i, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
}

func Sort(s sort.Interface) {
	length := s.Len()
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s.Less(j, i) {
				minIndex = j
			}
		}
		s.Swap(minIndex, i)
	}
}

func SelectionSort(arr []int, length int) {
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[j-1] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func (t Teacher) String() string {
	return fmt.Sprintf("Teacher: %s %v", t.Name, t.Score)
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(arr, len(arr))

	fmt.Println(arr)

	teachers := Teachers{Teacher{Name: "A", Score: 90}, Teacher{Name: "B", Score: 40}, Teacher{Name: "C", Score: 30}, Teacher{Name: "D", Score: 100}}
	Sort(teachers)

	for _, teacher := range teachers {
		fmt.Println(teacher)
	}
}
