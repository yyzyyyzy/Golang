// 在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，
// 然后在每一个有常量声明的行加一。
package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const (
		d = iota //再次定义常量时，iota会被重置为0
		e = iota
	)
	fmt.Println(d, e)

	const (
		f = iota //一个括号内可以只写一个iota，减少代码量
		g
		h
	)
	fmt.Println(f, g, h)

	const (
		o       = iota
		x, y, z = iota, iota, iota //如果多变量定义为iota，值相同
		p       = iota
	)
	fmt.Println(o, x, y, z, p)
}
