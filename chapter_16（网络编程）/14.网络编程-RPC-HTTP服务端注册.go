package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type PlusStruct struct {
	num1 int
	num2 int
}

func (p PlusStruct) Plus(plusStruct PlusStruct, result *int) (err error) {
	*result = plusStruct.num1 - plusStruct.num2
	return nil
}
func main() {
	plusStruct := PlusStruct{
		30,
		20,
	}
	rpc.RegisterName("PlusService", plusStruct)
	http.HandleFunc("/plus", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	})
	http.ListenAndServe(":8080", nil)
}
