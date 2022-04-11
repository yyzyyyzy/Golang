package main

import (
	"fmt"
	"runtime"
)

// 声明描述错误的结构体，成员是保存错误的执行函数
type panicContext struct {
	function string
}

// 创建一个保护函数ProtectRun
func ProtectRun(entry func()) {
	defer func() { // 延迟处理函数
		err := recover() // 发生宕机时，recover()获取panic传入的上下文并打印。defer延迟调用闭包

		switch err.(type) { // 使用switch对变量err进行类型断言
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	entry() // 对入参函数进行调用，当发生panic时，先调用defer后的闭包，再退出
}

func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手动宕机前")

		panic(&panicContext{ // 使用panic手动触发一个错误，并把一个结构体附带信息传入，被ProtectRun函数的recover()捕获
			"手动触发panic",
		})

		fmt.Println("手动宕机后") // panic之后的代码不再执行
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 1 // 模拟代码中空指针赋值造成的错误，此时Runtime层抛出错误，被ProtectRun函数的reover()捕获

		fmt.Println("赋值宕机后") // panic之后的代码不再执行
	})

	fmt.Println("运行后")
}
