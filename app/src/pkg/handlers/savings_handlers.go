package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func AddSavings(w http.ResponseWriter, r *http.Request) {

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

}

func EditSavings(w http.ResponseWriter, r *http.Request) {

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
}

func DeleteSavings(w http.ResponseWriter, r *http.Request) {

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
}
