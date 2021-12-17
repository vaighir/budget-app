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

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	fundName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert fund amount to float")
		app.Session.Put(r.Context(), "warning", "Fund must be a float.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	fund := models.Fund{
		Name:   fundName,
		Amount: amount,
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

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	fundName := r.Form.Get("name")
	amountAsString := r.Form.Get("amount")
	fundIdAsString := r.Form.Get("fund_id")

	amount, err := strconv.ParseFloat(amountAsString, 64)
	if err != nil {
		log.Println("Failed to convert fund amount to float")
		app.Session.Put(r.Context(), "warning", "Fund must be a float.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	fundId, err := strconv.Atoi(fundIdAsString)
	if err != nil {
		log.Println("Failed to convert fund id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't edit fund.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	fund := db_helpers.GetFundById(fundId)

	if redirectWrongHousehold(w, r, fund.HouseholdId, householdId, "edit", "a fund") {
		return
	}

	fund.Name = fundName
	fund.Amount = amount

	db_helpers.UpdateFund(fund)
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

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	fundIdAsString := r.Form.Get("fund_id")

	fundId, err := strconv.Atoi(fundIdAsString)
	if err != nil {
		log.Println("Failed to convert fund id to int")
		app.Session.Put(r.Context(), "warning", "Couldn't delete fund.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	fund := db_helpers.GetFundById(fundId)

	if redirectWrongHousehold(w, r, fund.HouseholdId, householdId, "delete", "a fund") {
		return
	}

	db_helpers.DeleteFund(fundId)
	app.Session.Put(r.Context(), "info", "Deleted fund.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}
