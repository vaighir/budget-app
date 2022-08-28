package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/vaighir/budget-app/app/pkg/config"
	"github.com/vaighir/budget-app/app/pkg/handlers"
	"github.com/vaighir/budget-app/app/pkg/helpers"
	"github.com/vaighir/budget-app/app/pkg/render"
)

const (
	portNumber        = ":8080"
	sessionLifetime   = 24 * time.Hour
	cookie_persist    = true
	minUsernameLength = 5
	maxUsernameLength = 50
	minPasswordLength = 10
)

var inProd = false

var app config.AppConfig
var session *scs.SessionManager

func main() {

	setupConfig()

	initialize()

	startServer()
}

func startServer() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
	log.Printf("Serving app on port %s", portNumber)
}

func setupConfig() {
	app.InProduction = inProd

	app.MinPasswordLength = minPasswordLength
	app.MinUsernameLength = minUsernameLength
	app.MaxUsernameLength = maxUsernameLength

	session = scs.New()
	session.Lifetime = sessionLifetime
	session.Cookie.Persist = cookie_persist
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		panic("cannot create template cache")
	}

	app.TemplateCache = templateCache
}

func initialize() {

	render.InitializeTemplates(&app)
	handlers.InitializeHandlers(&app)
	helpers.InitializeHelpers(&app)
}
