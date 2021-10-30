package handlers

import (
	"fmt"
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/config"
	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
	"github.com/vaighir/go-diet/app/pkg/render"
)

var app *config.AppConfig

func InitializeHandlers(a *config.AppConfig) {
	app = a
}

func Home(w http.ResponseWriter, r *http.Request) {

	app.Session.Put(r.Context(), "username", "admin")

	var user = db_helpers.GetUserById(1)

	stringMap := make(map[string]string)
	stringMap["username"] = user.Username

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func ShowUser(w http.ResponseWriter, r *http.Request) {

	var username = app.Session.Get(r.Context(), "username")

	var user = db_helpers.GetUserByUsername(fmt.Sprint(username))

	stringMap := make(map[string]string)
	stringMap["username"] = user.Username

	render.RenderTemplate(w, "show_user.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
