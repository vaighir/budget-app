package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vaighir/budget-app/app/pkg/config"
	"github.com/vaighir/budget-app/app/pkg/handlers"
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

	// Household routes

	mux.Get("/household", handlers.Household)

	mux.Get("/create-a-household", handlers.ShowNewHouseholdForm)
	mux.Post("/create-a-household", handlers.AddHousehold)

	// Income routes

	mux.Post("/add-income", handlers.AddIncome)
	mux.Post("/delete-income", handlers.DeleteIncome)
	mux.Post("/edit-income", handlers.EditIncome)

	// Savings routes

	mux.Post("/add-savings", handlers.AddSavings)
	mux.Post("/delete-savings", handlers.DeleteSavings)
	mux.Post("/edit-savings", handlers.EditSavings)

	// Savings routes

	mux.Post("/add-fund", handlers.AddFund)
	mux.Post("/delete-fund", handlers.DeleteFund)
	mux.Post("/edit-fund", handlers.EditFund)

	// Load static files

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
