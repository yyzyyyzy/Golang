package main

import "net/http"

func main() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Home"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
