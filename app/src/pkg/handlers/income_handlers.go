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

	income, err := parseAddIncomeForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
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

	newIncome, oldIncomeId, err := parseEditIncomeForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	oldIncome := db_helpers.GetIncomeById(oldIncomeId)

	if redirectWrongHousehold(w, r, oldIncome.HouseholdId, householdId, "edit", "an income") {
		return
	}

	oldIncome.Name = newIncome.Name
	oldIncome.Amount = newIncome.Amount

	db_helpers.UpdateIncome(oldIncome)
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

	income, err := parseDeleteIncomeForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	if redirectWrongHousehold(w, r, income.HouseholdId, householdId, "delete", "an income") {
		return
	}

	db_helpers.DeleteIncome(income.Id)
	app.Session.Put(r.Context(), "info", "Deleted income.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func parseAddIncomeForm(r *http.Request) (models.Income, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Income{}, err
	}

	incomeName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert income amount to float")
		app.Session.Put(r.Context(), "warning", "Income must be a float.")
		return models.Income{}, err
	}

	income := models.Income{
		Name:   incomeName,
		Amount: amount,
	}

	return income, err
}

func parseEditIncomeForm(r *http.Request) (models.Income, int, error) {

	err := r.ParseForm()
	if err != nil {
		return models.Income{}, -1, err
	}

	incomeName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")

	incomeIdAsString := r.Form.Get("income_id")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert income amount to float")
		app.Session.Put(r.Context(), "warning", "Income must be a float.")
		return models.Income{}, -1, err
	}

	incomeId, err := strconv.Atoi(incomeIdAsString)
	if err != nil {
		log.Println("Failed to convert income id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't edit income.")
		return models.Income{}, -1, err
	}

	newIncome := models.Income{
		Name:   incomeName,
		Amount: amount,
	}

	return newIncome, incomeId, err

}

func parseDeleteIncomeForm(r *http.Request) (models.Income, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Income{}, err
	}

	incomeIdAsString := r.Form.Get("income_id")

	incomeId, err := strconv.Atoi(incomeIdAsString)
	if err != nil {
		log.Println("Failed to convert income id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't delete income.")
		return models.Income{}, err
	}

	income := db_helpers.GetIncomeById(incomeId)

	return income, err
}
