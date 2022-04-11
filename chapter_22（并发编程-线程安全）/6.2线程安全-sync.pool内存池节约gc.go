package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//sync.pool是golang标准库中提供的一个通用的数据结构，可以创建池化的对象。把不用的对象回收，避免被清理，使用的时候直接取用。
//sync.pool主要有两个方法Get和Put,以及一个New字段。
//New的字段类型是函数func() interface{},可以用来创建临时对象。
//Get方法会从池中取走一个元素，返回一个intrtface{}类型，也有可能返回nil（New字段未设置，有没有空闲的元素返回）
//Put方法会接受一个interface{}类型的元素，将元素返回给Pool，如果是nil，则会忽略
//go 1.13版本中引入了victim cache,会将pool内数据拷贝一份,避免GC将其清空,即使没有引用的内容也可以保留最多两轮GC

var pool = sync.Pool{
	New: func() interface{} {
		return "123"
	},
}

func main() {
	t := pool.Get().(string)
	fmt.Println(t)

	pool.Put("321")
	pool.Put("321")

	runtime.GC()
	time.Sleep(1 * time.Second)

	t2 := pool.Get().(string)
	fmt.Println(t2)

	runtime.GC()
	time.Sleep(1 * time.Second)

	t2 = pool.Get().(string)
	fmt.Println(t2)
}
