package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	HandleError(err, "os.OpenFile")
	defer file.Close()
	return file.Write(data)
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "go: ", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil || len(msg) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, err)
		os.Exit(1)
	}
}

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
		log.Println(server.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v服务启动了，按任意键停止服务", serviceName)
		var s string
		fmt.Scanln(&s)
		server.Shutdown(ctx)
		cancel()
	}()

	return ctx
}

func main() {
	//创建log文件
	Run("E:\\golandlearning\\chapter_28（分布式系统搭建）\\分布式日志.log")
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

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service "+"response with code %v", res.StatusCode)
	}
	return nil
}
