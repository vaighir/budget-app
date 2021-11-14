package handlers

import (
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
	"github.com/vaighir/go-diet/app/pkg/render"
)

func Households(w http.ResponseWriter, r *http.Request) {
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

		render.RenderTemplate(w, "households.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			IntMap:    intMap,
			BoolMap:   boolMap,
		})

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to view households")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
