package main

import (
	"fmt"
)

type Search interface {
	Search(word string) []string
}

type Baidu_method1 struct {
}

func (B Baidu_method1) Search(word string) []string {
	result := []string{"评分10分-请开启vip获得权限查看", "评分5分-请开启vip获得权限查看", "评分1分-请开启vip获得权限查看"}
	fmt.Println(word, result)
	return result
}

type Baidu_method2 struct {
}

func (B Baidu_method2) Search(word string) []string {
	result := []string{"评分10分-请开启vip获得权限查看", "评分5分-请开启vip获得权限查看", "评分1分-三笑才子佳人"}
	fmt.Println(word, result)
	return result
}

type Baidu_method3 struct {
}

func (B Baidu_method3) Search(word string) []string {
	result := []string{"评分10分-肖生客的救赎", "评分5分-穿越是空的爱恋", "评分1分-三笑才子佳人"}
	fmt.Println(word, result)
	return result
}

func Baiducando(s Search) {
	s.Search("好电影")
}
func main() {
	B1 := Baidu_method1{}
	B2 := Baidu_method2{}
	B3 := Baidu_method3{}
	Baiducando(B1)
	Baiducando(B2)
	Baiducando(B3)
}
