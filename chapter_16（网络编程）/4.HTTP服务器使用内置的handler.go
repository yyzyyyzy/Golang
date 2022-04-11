package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "E:\\golandlearning\\chapter_24（Gin框架）\\fronthtml.html")
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
