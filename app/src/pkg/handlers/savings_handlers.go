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
	householdId := getHouseholdId(r)

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
	householdId := getHouseholdId(r)

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

	if redirectWrongHousehold(w, r, savings.HouseholdId, householdId, "edit", "savings") {
		return
	}

	savings.Name = savingsName
	savings.Amount = amount

	db_helpers.UpdateSavings(savings)
	app.Session.Put(r.Context(), "info", "Edited savings.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

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
	householdId := getHouseholdId(r)

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

	if redirectWrongHousehold(w, r, savings.HouseholdId, householdId, "delete", "savings") {
		return
	}

	db_helpers.DeleteSavings(savingsId)
	app.Session.Put(r.Context(), "info", "Deleted savings.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)
}
