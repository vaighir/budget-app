package handlers

import (
	"fmt"
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

func redirectIfNotLoggedIn(w http.ResponseWriter, r *http.Request, resource string) bool {
	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if !loggedIn {

		msg := fmt.Sprintf("You have to be logged in to view %s", resource)
		app.Session.Put(r.Context(), "warning", msg)
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return true
	} else {
		return false
	}
}

func redirectIfNoHousehold(w http.ResponseWriter, r *http.Request) bool {
	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	if householdId > 0 {
		return false
	} else {
		app.Session.Put(r.Context(), "warning", "You don't have a household.")
		http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
		return true
	}
}

func redirectIfHouseholdExists(w http.ResponseWriter, r *http.Request) bool {
	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	if householdId > 0 {
		app.Session.Put(r.Context(), "warning", "You already have a household.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return true
	} else {
		return false
	}
}

func redirectWrongHousehold(w http.ResponseWriter, r *http.Request, resourceHouseholdId int, householdId int, action string, resource string) bool {
	if resourceHouseholdId == householdId {
		return false
	} else {
		msg := fmt.Sprintf("You can't %s %s for a household different than yours", action, resource)
		app.Session.Put(r.Context(), "warning", msg)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return true
	}
}

func getHouseholdId(r *http.Request) int {
	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	return user.HouseholdId
}
