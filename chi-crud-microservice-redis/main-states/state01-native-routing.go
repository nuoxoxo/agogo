package main

import (
	"fmt"
	"net/http"
	_"github.com/go-chi/chi/v5"
	_"io/ioutil"
)

func main() {

	server := &http.Server{
		Addr:    ":10086",
		Handler: http.HandlerFunc(handler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err/listen & serve", err)
	}
}

func handler(writer http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	writer.Write([]byte("Hello, World!\n"))

	fmt.Println(req.Method, path)

	if req.Method != http.MethodPost && req.Method != http.MethodGet {
		fmt.Println("\t/Not POST/", req.Method, req.Method == http.MethodGet)
	}
	if len(req.URL.Query()) == 0 {
		fmt.Println("\t/empty query")
	}
	for k, v := range req.URL.Query() {
		fmt.Println("\t/item", k, v)
	}

	fmt.Println("\t/end \n")
}
