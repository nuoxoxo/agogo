package handlers

import (
	"fmt"
	"net/http"
)

type FooHandlers struct{}

func (fh *FooHandlers) Create(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/Create")
}

func (fh *FooHandlers) Get(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/GET (List)")
}

func (fh *FooHandlers) Update(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/Update:PUT")
}
