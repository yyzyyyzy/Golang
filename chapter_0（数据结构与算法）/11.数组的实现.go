package main

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //抓取第几个元素
	Set(index int, newval interface{}) error    //修改第几个数据
	insert(index int, newval interface{}) error //插入数据
	Append(newval interface{})                  //追加数据
	Delete(index int) error                     //删除数据
	Clear()                                     //清空数据
	String() string                             //返回字符串类型数据
}

type ArrayList struct {
	datastore []interface{} //数组存储的泛型（字符串、整数、实数）
	theSize   int           //数组的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList) //初始化结构体
	list.datastore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.datastore[index], nil
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.datastore[index] = newval
	return nil
}

func (list *ArrayList) insert(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}

	return nil
}

func (list *ArrayList) Append(newval interface{}) {
	list.datastore = append(list.datastore, newval) //叠加数据
	list.theSize++
}

func (list *ArrayList) Delete(index int) error {
	list.datastore = append(list.datastore[:index], list.datastore[index+1:]...)
	list.theSize--
	return nil
}

func (list *ArrayList) Clear() {
	list.datastore = make([]interface{}, 0, 10)
	list.theSize = 0
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.datastore)
}

func (list *ArrayList) checkisFull() {
	if list.theSize == cap(list.datastore) {
		newdataStore := make([]interface{}, 0, 2*list.theSize)
		copy(newdataStore, list.datastore)
		list.datastore = newdataStore
	}
}
func main() {
	arr := NewArrayList()
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)
	arr.Append(4)
	arr.Append(5)
	fmt.Println(arr)
	fmt.Println(arr.Set(1, "1090"))
	fmt.Println(arr.Get(1))
	fmt.Println(arr.String())

}
