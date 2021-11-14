package handlers

import (
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
	"github.com/vaighir/go-diet/app/pkg/render"
)

func Household(w http.ResponseWriter, r *http.Request) {
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
		} else {
			app.Session.Put(r.Context(), "warning", "You don't have a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

		stringMap["username"] = user.Username
		boolMap["logged_in"] = true

		render.RenderTemplate(w, "household.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			IntMap:    intMap,
			BoolMap:   boolMap,
		})

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to view households")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ShowNewHouseholdForm(w http.ResponseWriter, r *http.Request) {
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
			app.Session.Put(r.Context(), "warning", "You already have a household")
			http.Redirect(w, r, "/household", http.StatusSeeOther)

			return
		}

		stringMap["username"] = user.Username
		boolMap["logged_in"] = true

		render.RenderTemplate(w, "add-household-form.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			IntMap:    intMap,
			BoolMap:   boolMap,
		})

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to create a household")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func AddHousehold(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {
			app.Session.Put(r.Context(), "warning", "You already have a household")
			http.Redirect(w, r, "/household", http.StatusSeeOther)

			return
		}

		// TODO parse the form and create a household

		app.Session.Put(r.Context(), "info", "Household created")
		http.Redirect(w, r, "/household", http.StatusSeeOther)

		return

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to create a household")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
