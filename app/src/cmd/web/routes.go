package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vaighir/go-diet/app/pkg/config"
	"github.com/vaighir/go-diet/app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	// Base routes

	mux.Get("/", handlers.Home)

	// Auth routes

	mux.Get("/register", handlers.ShowRegisterForm)
	mux.Post("/register", handlers.Register)

	mux.Get("/login", handlers.ShowLoginForm)
	mux.Post("/login", handlers.Login)

	mux.Get("/logout", handlers.Logout)

	// Load static files

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
