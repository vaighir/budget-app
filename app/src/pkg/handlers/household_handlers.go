package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
	"github.com/vaighir/budget-app/app/pkg/render"
)

func Household(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "household") {
		return
	}

	if redirectIfNoHousehold(w, r) {
		return
	}

	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)
	intMap := make(map[string]int)
	floatMap := make(map[string]float64)
	interfaceMap := make(map[string]interface{})

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))
	householdId := user.HouseholdId

	household := db_helpers.GetHouseholdById(householdId)

	householdIncomes := db_helpers.GetAllIncomesByHouseholdId(householdId)

	var totalHouseholdIncome float64

	for _, income := range householdIncomes {
		totalHouseholdIncome += income.Amount
	}

	householdSavings := db_helpers.GetAllSavingsByHouseholdId(householdId)

	var totalHouseholdSavings float64

	for _, savings := range householdSavings {
		totalHouseholdSavings += savings.Amount
	}

	householdFunds := db_helpers.GetAllFundsByHouseholdId(householdId)

	var totalHouseholdFunds float64

	for _, fund := range householdFunds {
		totalHouseholdFunds += fund.Amount
	}

	householdMExpenses := db_helpers.GetAllMonthlyExpensesByHouseholdId(householdId)

	var totalHouseholdMExpenses float64

	for _, mExpense := range householdMExpenses {
		totalHouseholdMExpenses += mExpense.Amount
	}

	householdUExpenses := db_helpers.GetAllUpcomingExpensesByHouseholdId(householdId)

	var totalHouseholdUExpenses float64

	for _, uExpense := range householdUExpenses {
		totalHouseholdUExpenses += uExpense.Amount
	}

	emergencyFundAmount := float64(household.MonthsOfEmergencyFund) * totalHouseholdMExpenses

	totalHouseholdFunds += emergencyFundAmount

	monthlyBalance := totalHouseholdIncome - totalHouseholdMExpenses

	stringMap["username"] = user.Username
	stringMap["household_name"] = household.Name
	stringMap["picked-date"] = app.Session.PopString(r.Context(), "picked-date")

	boolMap["logged_in"] = true

	intMap["emergency_fund_length"] = household.MonthsOfEmergencyFund
	intMap["household_id"] = householdId

	floatMap["total_income"] = totalHouseholdIncome
	floatMap["total_savings"] = totalHouseholdSavings
	floatMap["total_funds"] = totalHouseholdFunds
	floatMap["total_monthly_expenses"] = totalHouseholdMExpenses
	floatMap["total_upcoming_expenses"] = totalHouseholdUExpenses
	floatMap["monthly_balance"] = monthlyBalance
	floatMap["emergency_fund_amount"] = emergencyFundAmount

	interfaceMap["incomes"] = householdIncomes
	interfaceMap["savings"] = householdSavings
	interfaceMap["funds"] = householdFunds
	interfaceMap["monthly_expenses"] = householdMExpenses
	interfaceMap["upcoming_expenses"] = householdUExpenses

	render.RenderTemplate(w, "household.page.tmpl", &models.TemplateData{
		StringMap:    stringMap,
		IntMap:       intMap,
		FloatMap:     floatMap,
		BoolMap:      boolMap,
		InterfaceMap: interfaceMap,
	})
}

func ChangeEmergencyFundLength(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "emergency fund") {
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

	household := db_helpers.GetHouseholdById(householdId)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	// TODO get household id from form and compare to see if the user is trying to change the right emergency fund
	emergencyFundLengthAsString := r.Form.Get("length")

	emergencyFundLength, err := strconv.Atoi(emergencyFundLengthAsString)
	if err != nil {
		log.Println("Failed to convert emergency fund length to int")
		log.Println(err)
		app.Session.Put(r.Context(), "warning", "Couldn't change emergency fund length.")
		http.Redirect(w, r, "/household", http.StatusSeeOther)
		return
	}

	household.MonthsOfEmergencyFund = emergencyFundLength

	db_helpers.UpdateHousehold(household)

	app.Session.Put(r.Context(), "info", "Updated emergency fund.")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

}

func ShowNewHouseholdForm(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "emergency fund") {
		return
	}

	if redirectIfHouseholdExists(w, r) {
		return
	}

	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)
	intMap := make(map[string]int)

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))

	stringMap["username"] = user.Username
	boolMap["logged_in"] = true

	render.RenderTemplate(w, "add-household-form.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IntMap:    intMap,
		BoolMap:   boolMap,
	})

}

// Parse data from create household form and create a household
func AddHousehold(w http.ResponseWriter, r *http.Request) {

	if redirectIfNotLoggedIn(w, r, "emergency fund") {
		return
	}

	if redirectIfHouseholdExists(w, r) {
		return
	}

	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	uid := app.Session.Get(r.Context(), "user_id")
	user := db_helpers.GetUserById(uid.(int))

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	householdName := r.Form.Get("name")

	household := models.Household{
		Name: householdName,
	}

	newHouseholdId := db_helpers.CreateHousehold(household)
	db_helpers.AddHouseholdToUser(user, newHouseholdId)

	app.Session.Put(r.Context(), "info", "Household created")
	http.Redirect(w, r, "/household", http.StatusSeeOther)

	log.Printf("Created household %s", householdName)

}

func DatePicker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	date := r.Form.Get("date")

	app.Session.Put(r.Context(), "picked-date", date)
	http.Redirect(w, r, "/household", http.StatusSeeOther)
}
