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

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))

		stringMap := make(map[string]string)
		stringMap["username"] = user.Username

		boolMap := make(map[string]bool)
		boolMap["logged_in"] = true

		render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			BoolMap:   boolMap,
		})
	} else {
		w.Write([]byte("You're not logged in"))
	}
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
