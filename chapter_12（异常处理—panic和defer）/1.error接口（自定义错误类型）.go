package main

import (
	"fmt"
)

type DivError struct {
	errstr string
}

func (De *DivError) Error() string {
	return De.errstr
}

/*
type error interface {
	Error() string
}
*/

func div1(dividend int, divisor int) (int, error) {
	//判断除数为0的情况并返回
	if divisor == 0 {
		return 0, &DivError{errstr: "被除数不为0"}
	}
	//正常计算，返回空错误
	return dividend / divisor, nil
}

func main() {
	fmt.Println(div1(24, 8)) //3 <nil>
	fmt.Println(div1(8, 0))  //0 自定义错误
}
