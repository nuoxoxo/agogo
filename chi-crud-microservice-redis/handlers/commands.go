package handlers

import (
	"fmt"
	"net/http"
)

type CommandHandlers struct{}

func (cmdh *CommandHandlers) Create(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/Create")
}

func (cmdh *CommandHandlers) List(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/List")
}

func (cmdh *CommandHandlers) GetByID(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/GetByID")
}

func (cmdh *CommandHandlers) UpdateByID(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/UpdateByID")
}

func (cmdh *CommandHandlers) DeleteByID(
	writer http.ResponseWriter,
	req *http.Request,
) {
	fmt.Println("/DeleteByID")
}
