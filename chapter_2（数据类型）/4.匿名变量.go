package main

import "fmt"

//我们在使用传统的强类型语言编程时，经常会出现这种情况，即在调用函数时为了获取一个值，
//却因为该函数返回多个值而不得不定义一堆没用的变量。
//在Go中这种情况可以通过结合使用多重返回和匿名变量来避免这种丑陋的写法，让代码看起来更加优雅。
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

//若只想获得nickName，则函数调用语句可以用如下方式编写：
func main() {
	_, _, nickName := GetName()
	fmt.Println("nickname=", nickName)
}
