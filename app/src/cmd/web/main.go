package main

import (
	"log"
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/config"
	"github.com/vaighir/go-diet/app/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	setupConfig()

	startServer()
}

func startServer() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving app on port %s", portNumber)
}

func setupConfig() {
	app.InProduction = false

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache

	render.NewTemplates(&app)

}
