package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}

//双指针
func isPalindrome(s string) bool {
	s = strings.ToLower(s) //大写转小写
	left, right := 0, len(s)-1
	for left < right {
		if !isalnum(s[left]) {
			left++
			continue
		}
		if !isalnum(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isalnum(ch byte) bool {
	// Ascii编码中 0-9:[48-57], A-Z:[65-90], a-z:[97-122], 把不属于数字、字母筛选去除
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
