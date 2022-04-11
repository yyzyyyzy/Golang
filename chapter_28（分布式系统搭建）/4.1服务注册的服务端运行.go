package main

import (
	"encoding/json"
	"net/http"
	"sync"
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
