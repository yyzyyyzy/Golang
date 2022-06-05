package main

import "fmt"

func main() {
	//右移运算符>>：低位溢出，符号位不变，并用符号位补溢出的高位，正数和负数补位的时候补的不一样，负数补1，正数补0
	//左移运算符<<：左移就是把所有位向左移动几位，低位补0
	//案例演示：
	//a := 1 >> 2 //0000 0001 => 0000 0000 = 0
	//c := 1 << 2 // 0000 0001 => 0000 0100 = 4

	var a int = 1 >> 2
	var b int = -1 >> 2
	var e int = -2 >> 2
	var c int = 1 << 2
	var d int = -1 << 2

	//a,b,c,d的结果是多少
	fmt.Println("a=", a) //0
	fmt.Println("b=", b) //-1
	fmt.Println("e=", e) //-1
	fmt.Println("c=", c) //4
	fmt.Println("d=", d) //-4
}
