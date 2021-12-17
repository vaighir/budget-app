package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func AddIncome(w http.ResponseWriter, r *http.Request) {

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

}

func EditIncome(w http.ResponseWriter, r *http.Request) {

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

	incomeName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")
	incomeIdAsString := r.Form.Get("income_id")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert income amount to float")
		app.Session.Put(r.Context(), "warning", "Income must be a float.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	incomeId, err := strconv.Atoi(incomeIdAsString)
	if err != nil {
		log.Println("Failed to convert income id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't edit income.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	income := db_helpers.GetIncomeById(incomeId)

	if redirectWrongHousehold(w, r, income.HouseholdId, householdId, "edit", "an income") {
		return
	}

	income.Name = incomeName
	income.Amount = amount

	db_helpers.UpdateIncome(income)
	app.Session.Put(r.Context(), "info", "Edited income.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func DeleteIncome(w http.ResponseWriter, r *http.Request) {

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

	incomeIdAsString := r.Form.Get("income_id")

	incomeId, err := strconv.Atoi(incomeIdAsString)
	if err != nil {
		log.Println("Failed to convert income id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't delete income.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	income := db_helpers.GetIncomeById(incomeId)

	if redirectWrongHousehold(w, r, income.HouseholdId, householdId, "delete", "an income") {
		return
	}

	db_helpers.DeleteIncome(incomeId)
	app.Session.Put(r.Context(), "info", "Deleted income.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}
