package main

import (
	"fmt"
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string //新类型的别名为fileLog

//写入数据的方法
func (fl fileLog) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	HandleError(err, "os.OpenFile")
	defer file.Close()
	return file.Write(data)
}

//把log指向某个文件地址
func Run(destination string) {
	log = stlog.New(fileLog(destination), "go", stlog.LstdFlags)
}

//路由/log使用post方法时，写入数据到log文件
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
