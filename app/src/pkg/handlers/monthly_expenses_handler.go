package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func AddMonthlyExpense(w http.ResponseWriter, r *http.Request) {
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

			mExpenseName := r.Form.Get("name")
			amountAsString := r.Form.Get("amount")

			amount, err := strconv.ParseFloat(amountAsString, 64)
			if err != nil {
				log.Println("Failed to convert monthly expense amount to float")
				app.Session.Put(r.Context(), "warning", "Monthly expense must be a float.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			mExpense := models.MonthlyExpense{
				Name:   mExpenseName,
				Amount: amount,
			}

			db_helpers.CreateMonthlyExpence(householdId, mExpense)

			app.Session.Put(r.Context(), "info", "Added monthly expense to your household.")
			http.Redirect(w, r, "/household", http.StatusSeeOther)
			return

		} else {
			app.Session.Put(r.Context(), "warning", "You can't add monthly expense without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to add an monthly expense")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditMonthlyExpense(w http.ResponseWriter, r *http.Request) {
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

			mExpenseName := r.Form.Get("name")
			amountAsString := r.Form.Get("amount")
			mExpenseIdAsString := r.Form.Get("monthly_expense_id")

			amount, err := strconv.ParseFloat(amountAsString, 64)
			if err != nil {
				log.Println("Failed to convert monthly expense amount to float")
				app.Session.Put(r.Context(), "warning", "Monthly expense must be a float.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			mExpenseId, err := strconv.Atoi(mExpenseIdAsString)
			if err != nil {
				log.Println("Failed to convert monthly expense id to int")
				app.Session.Put(r.Context(), "warning", "Couldn't edit monthly expense.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			mExpense := db_helpers.GetMonthlyExpenseById(mExpenseId)

			if mExpense.HouseholdId == householdId {

				mExpense.Name = mExpenseName
				mExpense.Amount = amount

				db_helpers.UpdateMonthlyExpense(mExpense)
				app.Session.Put(r.Context(), "info", "Edited monthly expense.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return

			} else {
				app.Session.Put(r.Context(), "warning", "You can't edit monthly expense for a different household than yours.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

		} else {
			app.Session.Put(r.Context(), "warning", "You can't edit monthly expense without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to edit monthly expense")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeleteMonthlyExpense(w http.ResponseWriter, r *http.Request) {
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

			mExpenseIdAsString := r.Form.Get("monthly_expense_id")

			mExpenseId, err := strconv.Atoi(mExpenseIdAsString)
			if err != nil {
				log.Println("Failed to convert monthly expense id to int")
				app.Session.Put(r.Context(), "warning", "Couldn't delete monthly expense.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			mExpense := db_helpers.GetMonthlyExpenseById(mExpenseId)

			if mExpense.HouseholdId == householdId {
				db_helpers.DeleteMonthlyExpense(mExpenseId)
				app.Session.Put(r.Context(), "info", "Deleted monthly expense.")
				http.Redirect(w, r, "/household", http.StatusSeeOther)
				return
			}

			app.Session.Put(r.Context(), "warning", "You can only delete monthly expense for your own household.")
			http.Redirect(w, r, "/household", http.StatusSeeOther)
			return

		} else {
			app.Session.Put(r.Context(), "warning", "You can't add monthly expense without a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to add monthly expense")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
