package main

import "net/http"

//自定义handler接口
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	mh := myHandler{}
	http.ListenAndServe("localhost:8080", &mh) //http.DefaultServeMux 是一个路由分发的handler接口
}
