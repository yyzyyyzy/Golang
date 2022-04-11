//单引号表示字符，双引号表示字符串
//字符类型（byte）uint8类型等价
package main

import "fmt"

func main() {
	var string1 byte = 'a'

	//输出的是对应的字符类型的ASCII码值
	fmt.Println("string1=", string1)

	//输出的是对应的字符类型的值
	var string2 byte = 'a'

	//var string3 byte = '李' //已经overflow并超出byte存储的范围
	var string3 int = '李'
	fmt.Printf("string2=%c \nstring3=%d", string2, string3)
}
