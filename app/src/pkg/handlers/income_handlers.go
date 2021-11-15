package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
)

func AddIncome(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {

			// TODO parse form, add income and redirect to household

			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}

			incomeName := r.Form.Get("name")
			amountAsString := r.Form.Get("amount")

			amount, err := strconv.ParseFloat(amountAsString, 64)
			if err != nil {
				log.Println("Failed to convert income amount to float")
				app.Session.Put(r.Context(), "warning", "Income must be a float.")
				http.Redirect(w, r, "/add-income", http.StatusSeeOther)
				return
			}

			income := models.Income{
				Name:   incomeName,
				Amount: amount,
			}

			db_helpers.CreateIncome(householdId, income)

			app.Session.Put(r.Context(), "info", "Added income to your household.")
			http.Redirect(w, r, "/household", http.StatusSeeOther)
			return

		} else {
			app.Session.Put(r.Context(), "warning", "You can't add income without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to add an income")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
