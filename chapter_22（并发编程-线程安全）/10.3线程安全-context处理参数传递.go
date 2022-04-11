package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("916990143", "123qqqAAA...")
}

func ProcessRequest(UserName, PassWord string) {
	ctx := context.WithValue(context.Background(), "UserName", UserName)
	ctx = context.WithValue(ctx, "PassWord", PassWord)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("用户名:%v 密码:%v", ctx.Value("UserName"), ctx.Value("PassWord"))
}

/*
使用的key必须是可比较的,也就是说== 和 != 必须能返回正确的结果
返回值必须是并发安全的,这样才能从多个goroutine访问
*/
