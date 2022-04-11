package main

import "net/http"

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type showHandler struct{}

func (s *showHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show your penis"))
}

func main() {
	mh := helloHandler{}
	sh := showHandler{}

	//2.使用http.handle将某个路由附加到DefaultServeMux
	http.Handle("/hello", &mh)
	http.Handle("/show", &sh)

	//1.使用默认的路由分发接口 nil=DefaultServeMux
	http.ListenAndServe("localhost:8080", nil)
}
