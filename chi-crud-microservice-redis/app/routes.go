package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get(
		"/",
		func(writer http.ResponseWriter, r *http.Request) {
			writer.WriteHeader(http.StatusOK)
		},
	)
	return router
}
