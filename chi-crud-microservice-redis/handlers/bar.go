package handlers

import (
	"fmt"
	"net/http"
)

type BarHandlers struct{}

func (bh *BarHandlers) Create(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/Create")
}

func (bh *BarHandlers) Get(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/GET (List)")
}

func (bh *BarHandlers) Update(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/Update:PATCH")
}
