package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

const (
	LogService = ServiceName("LogService")
)

const (
	ServerPort = ":3000"
	ServiceURL = "http://localhost" + ServerPort + "/services"
)

type registry struct {
	Registration []Registration //已经注册的服务（可能被并发访问）
	mutux        *sync.Mutex    //保证线程安全
}

func (r *registry) add(reg Registration) error {
	r.mutux.Lock()
	r.Registration = append(r.Registration, reg)
	r.mutux.Unlock()
	return nil
}

var reg = registry{
	Registration: make([]Registration, 0),
	mutux:        new(sync.Mutex),
}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("Request received")
	switch request.Method {
	case http.MethodPost:
		dec := json.NewDecoder(request.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

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
