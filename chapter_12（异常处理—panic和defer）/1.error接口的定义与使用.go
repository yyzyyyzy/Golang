package main

import (
	"errors" //需要导入errors包
	"fmt"
)

// 定义一个除数为0的错误对象
var errDivisionByZero = errors.New("除数为0")

func div(dividend int, divisor int) (int, error) {
	//判断除数为0的情况并返回
	if divisor == 0 {
		return 0, errDivisionByZero
	}
	//正常计算，返回空错误
	return dividend / divisor, nil
}

func main() {
	fmt.Println(div(24, 8)) //3 <nil>
	fmt.Println(div(8, 0))  //0 division by zero
}
