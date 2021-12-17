package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func AddFund(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)
	householdId := getHouseholdId(r)

	fund, err := parseAddFundForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	db_helpers.CreateFund(householdId, fund)

	app.Session.Put(r.Context(), "info", "Added fund to your household.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func EditFund(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)
	householdId := getHouseholdId(r)

	newFund, oldFundId, err := parseEditFundForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	oldFund := db_helpers.GetFundById(oldFundId)

	if redirectWrongHousehold(w, r, oldFund.HouseholdId, householdId, "edit", "a fund") {
		return
	}

	oldFund.Name = newFund.Name
	oldFund.Amount = newFund.Amount

	db_helpers.UpdateFund(oldFund)
	app.Session.Put(r.Context(), "info", "Edited fund.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func DeleteFund(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)
	householdId := getHouseholdId(r)

	fund, err := parseDeleteFundForm(r)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	if redirectWrongHousehold(w, r, fund.HouseholdId, householdId, "delete", "a fund") {
		return
	}

	db_helpers.DeleteFund(fund.Id)
	app.Session.Put(r.Context(), "info", "Deleted fund.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func parseAddFundForm(r *http.Request) (models.Fund, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Fund{}, err
	}

	fundName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert fund amount to float")
		app.Session.Put(r.Context(), "warning", "Fund must be a float.")
		return models.Fund{}, err
	}

	fund := models.Fund{
		Name:   fundName,
		Amount: amount,
	}

	return fund, err
}

func parseEditFundForm(r *http.Request) (models.Fund, int, error) {

	err := r.ParseForm()
	if err != nil {
		return models.Fund{}, -1, err
	}

	fundName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")

	fundIdAsString := r.Form.Get("fund_id")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert fund amount to float")
		app.Session.Put(r.Context(), "warning", "Fund must be a float.")
		return models.Fund{}, -1, err
	}

	fundId, err := strconv.Atoi(fundIdAsString)
	if err != nil {
		log.Println("Failed to convert fund id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't edit fund.")
		return models.Fund{}, -1, err
	}

	newFund := models.Fund{
		Name:   fundName,
		Amount: amount,
	}

	return newFund, fundId, err

}

func parseDeleteFundForm(r *http.Request) (models.Fund, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Fund{}, err
	}

	fundIdAsString := r.Form.Get("fund_id")

	fundId, err := strconv.Atoi(fundIdAsString)
	if err != nil {
		log.Println("Failed to convert fund id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't delete fund.")
		return models.Fund{}, err
	}

	fund := db_helpers.GetFundById(fundId)

	return fund, err
}
