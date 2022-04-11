package main

import "fmt"

type Doctor struct {
	Name string
	Sex  string
}

//类型断言
//一个判断传入参数类型的函数
func judge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool,value is %v\n", index, v)
		case int, int64, int32:
			fmt.Printf("%d params is int,value is %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float,value is %v\n", index, v)
		case string:
			fmt.Printf("%d params is string,value is %v\n", index, v)
		case Doctor:
			fmt.Printf("%d params doctor,value is %v\n", index, v)
		case *Doctor:
			fmt.Printf("%d params *doctor,value is %v\n", index, v)

		}
	}
}
func main() {
	d := Doctor{
		Name: "doctor1",
		Sex:  "male",
	}
	judge(28, 8.2, "this is a test", d, &d)

}
