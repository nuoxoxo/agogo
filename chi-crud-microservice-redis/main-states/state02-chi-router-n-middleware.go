package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "io/ioutil"
	"net/http"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routes := []struct {
		Method string
		Path   string
	}{
		{Method: http.MethodGet, Path: "/"}, // GET
		{Method: http.MethodGet, Path: "/bar"},
		{Method: http.MethodGet, Path: "/foo"},
		{Method: http.MethodPost, Path: "/bar"}, // POST
		{Method: http.MethodPost, Path: "/foo"},
		{Method: http.MethodPatch, Path: "/bar"}, // PATCH
		{Method: http.MethodPut, Path: "/foo"},   // PUT
	}

	/*
		router := chi.NewRouter ()
		router.Get("/", handler)
		router.Get("/bar", handler)
		router.Get("/foo", handler)
		router.Post("/bar", handler)
		router.Post("/foo", handler)
		router.Patch("/bar", handler)
		router.Put("/foo", handler)
	*/

	for _, route := range routes {
		router.Method(
			route.Method,
			route.Path,
			http.HandlerFunc(handler),
		)
	}

	server := &http.Server{
		Addr:    ":10086",
		Handler: router, //http.HandlerFunc(handler),
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

	fmt.Print("\t/end \n\n")
}
