package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func AddSavings(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {

			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}

			savingsName := r.Form.Get("name")
			amountAsString := r.Form.Get("amount")

			amount, err := strconv.ParseFloat(amountAsString, 64)
			if err != nil {
				log.Println("Failed to convert savings amount to float")
				app.Session.Put(r.Context(), "warning", "Savings must be a float.")
				http.Redirect(w, r, "/add-savings", http.StatusSeeOther)
				return
			}

			savings := models.Savings{
				Name:   savingsName,
				Amount: amount,
			}

			db_helpers.CreateSavings(householdId, savings)

			app.Session.Put(r.Context(), "info", "Added savings to your household.")
			http.Redirect(w, r, "/household", http.StatusSeeOther)
			return

		} else {
			app.Session.Put(r.Context(), "warning", "You can't add savings without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to add an savings")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditSavings(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {

			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}

			savingsName := r.Form.Get("name")
			amountAsString := r.Form.Get("amount")
			savingsIdAsString := r.Form.Get("savings_id")

			amount, err := strconv.ParseFloat(amountAsString, 64)
			if err != nil {
				log.Println("Failed to convert savings amount to float")
				app.Session.Put(r.Context(), "warning", "Savings must be a float.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			savingsId, err := strconv.Atoi(savingsIdAsString)
			if err != nil {
				log.Println("Failed to convert savings id to int")
				app.Session.Put(r.Context(), "warning", "Couldn't edit savings.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			savings := db_helpers.GetSavingsById(savingsId)

			if savings.HouseholdId == householdId {

				savings.Name = savingsName
				savings.Amount = amount

				db_helpers.UpdateSavings(savings)
				app.Session.Put(r.Context(), "info", "Edited savings.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return

			} else {
				app.Session.Put(r.Context(), "warning", "You can't edit savings for a different household than yours.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

		} else {
			app.Session.Put(r.Context(), "warning", "You can't edit savings without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to edit savings")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeleteSavings(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {

			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}

			savingsIdAsString := r.Form.Get("savings_id")

			savingsId, err := strconv.Atoi(savingsIdAsString)
			if err != nil {
				log.Println("Failed to convert savings id to int")
				app.Session.Put(r.Context(), "warning", "Couldn't delete savings.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			savings := db_helpers.GetSavingsById(savingsId)

			if savings.HouseholdId == householdId {
				db_helpers.DeleteSavings(savingsId)
				app.Session.Put(r.Context(), "info", "Deleted savings.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			app.Session.Put(r.Context(), "warning", "You can only delete savings for your own household.")
			http.Redirect(w, r, "/household", http.StatusSeeOther)
			return

		} else {
			app.Session.Put(r.Context(), "warning", "You can't add savings without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to add savings")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
