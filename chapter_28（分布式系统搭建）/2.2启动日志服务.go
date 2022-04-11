package main

import (
	"context"
	"fmt"
	stlog "log"
)

func main() {
	Run("./distributed.log")
	host, port := "localhost", "4000"
	r := Registration{
		ServiceName: "Log Service",
		ServiceURL:  fmt.Sprintf("http://%s:%s", host, port),
	}
	ctx, err := Start(context.Background(), host, port, r, RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()

	fmt.Println("服务已关闭")
}
