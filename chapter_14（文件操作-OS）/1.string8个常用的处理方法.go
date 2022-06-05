package main

import (
	"fmt"
	"strings"
)

func main() {
	//查找 strings.Contains 返回bool
	fmt.Println("查找子串是否在指定的字符串中")
	fmt.Println(" Contains 函数的用法")
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true

	//查找 strings.Index 返回int
	fmt.Println("")
	fmt.Println(" Index 函数的用法")
	fmt.Println(strings.Index("NLT_abc", "abc")) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.Index("NLT_abc", "aaa")) // 子集没有在父级中出现过 返回-1
	fmt.Println(strings.Index("我是中国人", "中"))     // 中文字符3个字节 6

	//组合 strings.Join 返回string
	fmt.Println("")
	fmt.Println(" Join 函数的用法")
	slice := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(slice, ", ")) // 返回字符串：foo, bar, baz

	//重复 strings.Repeat 返回string
	fmt.Println("")
	fmt.Println(" Repeat 函数的用法")
	fmt.Println("ba" + strings.Repeat("na", 2)) // 字符串重复

	//替换 strings.Replace 返回string
	fmt.Println("")
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("性感网友在线性感吃屎", "性感", "**", 2))        // aBaACEDF
	fmt.Println(strings.Replace("性感网友性感取名性感吃屎性感大炮", "性感", "**", -1)) // 第四个参数小于0，表示所有的都替换

	//分割 strings.Split 返回string_slice
	fmt.Println("")
	fmt.Println(" Split 函数的用法")
	fmt.Println(strings.Split("136-108-50940", "-"))
	fmt.Println(strings.Split("916990143@qq.com", "@"))

	//去除冗余 strings.Trim 返回string
	//切割前缀使用TrimPrefix，切割后缀使用TrimSuffix
	fmt.Println("")
	fmt.Println(" Trim  函数的用法")
	fmt.Println(strings.Trim(" !!! Achtung !!! ", "! ")) // 去除字符串头尾不需要的内容

	//去除冗余 strings.Fields 返回slice
	fmt.Println("")
	fmt.Println(" Fields 函数的用法")
	fmt.Println(strings.Fields("  are        you        ok   ")) //去除多余空格

	//计数 strings.Count 返回int
	fmt.Println("")
	fmt.Println(" Count 函数的用法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现
}
