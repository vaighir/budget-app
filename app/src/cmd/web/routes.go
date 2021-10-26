package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vaighir/go-diet/app/pkg/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)

	return mux
}
