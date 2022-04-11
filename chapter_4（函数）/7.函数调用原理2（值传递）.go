package main

import "fmt"

func swap1(a, b *int) { //参数为指针，传参数就是从右往左拷贝参数的地址
	*a, *b = *b, *a //此处交换了指针所对应的值a=1,b=2 -> a=2,b=1
}

func main() {
	a, b := 1, 2
	swap1(&a, &b)
	fmt.Println(a, b)
}
