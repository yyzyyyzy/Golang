package main

import (
	"fmt"
)

//空接口可以接收任意类型，适合容器设计
type Dictionary struct {
	data map[interface{}]interface{} //key和value都是接口类型
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// 创建一个新字典
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{}) //新建相当于将原先的map交给gc，重新开辟一块内存空间存储新的map
}

func Newdictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {
	dict := Newdictionary()
	dict.Set(1, 10)
	dict.Set(2, 30)
	dict.Set(3, 50)
	dict.Set(4, "我叼")

	fmt.Println(dict.Get(3))
	dict.Visit(func(k, v interface{}) bool {
		value, ok := v.(int)
		if ok {
			if value >= 10 {
				fmt.Println(k)
				return true
			}
		}
		return true
	})
}
