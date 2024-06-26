package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", Home)


	return router
}