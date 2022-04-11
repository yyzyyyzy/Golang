package main

import "fmt"

//单例模式适用场景：（回收站、打印机、网站计数器。。。）
//				1.需要生成唯一序列的环境
//				2.需要频繁实例化然后销毁的对象。
//				3.创建对象时耗时过多或者耗资源过多，但又经常用到的对象。
//				4.方便资源相互通信的环境

type singleton struct{}

var instance *singleton

func (s singleton) PrintAddr() {
	fmt.Printf("%p\n", &s)
}

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	c1 := GetInstance()
	c2 := GetInstance()
	c3 := GetInstance()
	c1.PrintAddr()
	c2.PrintAddr()
	c3.PrintAddr()
}
