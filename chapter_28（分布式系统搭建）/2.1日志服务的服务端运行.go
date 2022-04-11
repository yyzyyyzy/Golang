package main

import (
	"context"
	"fmt"
	"net/http"
)

//启动服务
func Start(ctx context.Context, host, port string, reg Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var server http.Server
	server.Addr = ":" + port

	go func() {
		//服务启动发生错误，就结束服务，打印日志
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {
		//用户手动停止服务
		fmt.Printf("%v服务启动了，按任意键停止服务", serviceName)
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx) //关闭http.server
		cancel()
	}()

	return ctx
}
