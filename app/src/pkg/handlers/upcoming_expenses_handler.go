package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

var layout = "2006-01-02"

func AddUpcomingExpense(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uExpenseName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")
	deadlineAsString := r.Form.Get("deadline")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert upcoming expense amount to float")
		app.Session.Put(r.Context(), "warning", "Upcoming expense must be a float.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	deadline, err := time.Parse(layout, deadlineAsString)
	if err != nil {
		log.Printf("Failed to convert upcoming expense date: %s", deadlineAsString)
		app.Session.Put(r.Context(), "warning", "Failed to parse date.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	uExpense := models.UpcomingExpense{
		Name:     uExpenseName,
		Amount:   amount,
		Deadline: deadline,
	}

	db_helpers.CreateUpcomingExpence(householdId, uExpense)

	app.Session.Put(r.Context(), "info", "Added upcoming expense to your household.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func EditUpcomingExpense(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uExpenseName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")
	uExpenseIdAsString := r.Form.Get("upcoming_expense_id")
	deadlineAsString := r.Form.Get("deadline")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert upcoming expense amount to float")
		app.Session.Put(r.Context(), "warning", "Upcoming expense must be a float.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	deadline, err := time.Parse(layout, deadlineAsString)
	if err != nil {
		log.Printf("Failed to convert upcoming expense date: %s", deadlineAsString)
		app.Session.Put(r.Context(), "warning", "Failed to parse date.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	uExpenseId, err := strconv.Atoi(uExpenseIdAsString)
	if err != nil {
		log.Println("Failed to convert upcoming expense id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't edit upcoming expense.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	uExpense := db_helpers.GetUpcomingExpenseById(uExpenseId)

	if uExpense.HouseholdId == householdId {

		uExpense.Name = uExpenseName
		uExpense.Amount = amount
		uExpense.Deadline = deadline

		db_helpers.UpdateUpcomingExpense(uExpense)
		app.Session.Put(r.Context(), "info", "Edited upcoming expense.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return

	} else {
		app.Session.Put(r.Context(), "warning", "You can't edit upcoming expense for a different household than yours.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}
}

func DeleteUpcomingExpense(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uExpenseIdAsString := r.Form.Get("upcoming_expense_id")

	uExpenseId, err := strconv.Atoi(uExpenseIdAsString)
	if err != nil {
		log.Println("Failed to convert upcoming expense id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't delete upcoming expense.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	mExpense := db_helpers.GetUpcomingExpenseById(uExpenseId)

	if mExpense.HouseholdId == householdId {
		db_helpers.DeleteUpcomingExpense(uExpenseId)
		app.Session.Put(r.Context(), "info", "Deleted upcoming expense.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "warning", "You can only delete upcoming expense for your own household.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}
