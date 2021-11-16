package handlers

import (
	"log"
	"net/http"

	"github.com/vaighir/budget-app/app/pkg/db_helpers"
	"github.com/vaighir/budget-app/app/pkg/models"
	"github.com/vaighir/budget-app/app/pkg/render"
)

func Household(w http.ResponseWriter, r *http.Request) {
	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)
	intMap := make(map[string]int)
	floatMap := make(map[string]float64)
	interfaceMap := make(map[string]interface{})

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {
			intMap["household_id"] = householdId

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

			stringMap["username"] = user.Username
			stringMap["household_name"] = household.Name

			boolMap["logged_in"] = true

			intMap["emergency_fund"] = household.MonthsOfEmergencyFund

			floatMap["total_income"] = totalHouseholdIncome
			floatMap["total_savings"] = totalHouseholdSavings

			interfaceMap["incomes"] = householdIncomes
			interfaceMap["savings"] = householdSavings

			render.RenderTemplate(w, "household.page.tmpl", &models.TemplateData{
				StringMap:    stringMap,
				IntMap:       intMap,
				FloatMap:     floatMap,
				BoolMap:      boolMap,
				InterfaceMap: interfaceMap,
			})

			return

		} else {
			app.Session.Put(r.Context(), "warning", "You don't have a household.")
			http.Redirect(w, r, "/create-a-household", http.StatusSeeOther)
			return
		}

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to view households")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ShowNewHouseholdForm(w http.ResponseWriter, r *http.Request) {
	boolMap := make(map[string]bool)
	stringMap := make(map[string]string)
	intMap := make(map[string]int)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {
			app.Session.Put(r.Context(), "warning", "You already have a household")
			http.Redirect(w, r, "/household", http.StatusSeeOther)

			return
		}

		stringMap["username"] = user.Username
		boolMap["logged_in"] = true

		render.RenderTemplate(w, "add-household-form.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
			IntMap:    intMap,
			BoolMap:   boolMap,
		})

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to create a household")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Parse data from create household form and create a household
func AddHousehold(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	loggedIn := app.Session.Exists(r.Context(), "user_id")

	if loggedIn {
		uid := app.Session.Get(r.Context(), "user_id")
		user := db_helpers.GetUserById(uid.(int))
		householdId := user.HouseholdId

		if householdId > 0 {
			app.Session.Put(r.Context(), "warning", "You already have a household")
			http.Redirect(w, r, "/household", http.StatusSeeOther)

			return
		}

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

		return

	} else {
		app.Session.Put(r.Context(), "warning", "You have to be logged in to create a household")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
