package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}

func main() {
	router := httprouter.New()
	router.POST("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
