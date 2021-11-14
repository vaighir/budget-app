package handlers

import (
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

	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))

		stringMap["username"] = user.Username
		boolMap["logged_in"] = true

	} else {
		boolMap["logged_in"] = false
	}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		BoolMap:   boolMap,
	})

}

func getSessionMsg(r *http.Request, stringMap map[string]string) {
	stringMap["warning"] = app.Session.PopString(r.Context(), "warning")
	stringMap["info"] = app.Session.PopString(r.Context(), "info")
}
