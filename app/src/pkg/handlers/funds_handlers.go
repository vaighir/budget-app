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

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

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

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

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

	if fund.HouseholdId == householdId {

		fund.Name = fundName
		fund.Amount = amount

		db_helpers.UpdateFund(fund)
		app.Session.Put(r.Context(), "info", "Edited fund.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return

	} else {
		app.Session.Put(r.Context(), "warning", "You can't edit fund for a different household than yours.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}
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

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

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

	if fund.HouseholdId == householdId {
		db_helpers.DeleteFund(fundId)
		app.Session.Put(r.Context(), "info", "Deleted fund.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "warning", "You can only delete fund for your own household.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)
}
