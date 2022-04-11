package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/services", &RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		//用户手动停止服务
		fmt.Println("服务注册服务启动了，按任意键停止服务")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("服务注册服务已关闭")
}
