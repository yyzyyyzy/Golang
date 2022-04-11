package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	// 导入头文件 math/rand
	// 创建随机数种子
	// 创建随机数
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())            //生成比较大的随机数
	a, b := fmt.Println(rand.Intn(10)) //常用 取模10 生成0-9
	fmt.Printf("%T,%T", a, b)
	fmt.Println(reflect.TypeOf(a))
}
