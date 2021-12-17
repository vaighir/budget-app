package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vaighir/budget-app/app/pkg/config"
	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
	"github.com/vaighir/budget-app/app/pkg/render"
)

var app *config.AppConfig

func InitializeHandlers(a *config.AppConfig) {
	app = a
}

func Home(w http.ResponseWriter, r *http.Request) {

	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)
	intMap := make(map[string]int)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))

		householdId := user.HouseholdId

		if householdId > 0 {
			intMap["household_id"] = householdId
		}

		stringMap["username"] = user.Username
		boolMap["logged_in"] = true

	} else {
		boolMap["logged_in"] = false
	}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
		BoolMap:   boolMap,
	})

}

func getSessionMsg(r *http.Request, stringMap map[string]string) {
	stringMap["warning"] = app.Session.PopString(r.Context(), "warning")
	stringMap["info"] = app.Session.PopString(r.Context(), "info")
}

// TODO needs fixing
func redirectIfNotLoggedIn(w http.ResponseWriter, r *http.Request, resource string) bool {
	loggedIn := app.Session.Exists(r.Context(), "user_id")
	log.Printf("You're %t", loggedIn)

	if !loggedIn {

		msg := fmt.Sprintf("You have to be logged in to view %s", resource)
		log.Print(msg)

		app.Session.Put(r.Context(), "warning", msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return true
	} else {
		return false
	}
}
