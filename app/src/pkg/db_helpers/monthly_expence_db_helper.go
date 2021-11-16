package db_helpers

import (
	"log"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func CreateMonthlyExpence(householdId int, mExpese models.MonthlyExpense) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("insert into monthly_expences (household_id, name, amount) values ($1, $2, $3)", householdId, mExpese.Name, mExpese.Amount)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func GetAllMonthlyExpensesByHouseholdId(householdId int) []models.MonthlyExpense {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("select id, name, amount from monthly_expences where household_id = $1", householdId)
	if err != nil {
		log.Fatal(err)
	}

	var mExpese models.MonthlyExpense
	var mExpeses []models.MonthlyExpense

	for rows.Next() {
		err := rows.Scan(&mExpese.Id, &mExpese.Name, &mExpese.Amount)
		if err != nil {
			log.Println(err)
		}

		mExpeses = append(mExpeses, mExpese)

	}

	return mExpeses
}

func GetMonthlyExpenseById(id int) models.MonthlyExpense {

	var mExpense models.MonthlyExpense

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select household_id, name, amount from monthly_expences where id = $1", id).Scan(&mExpense.HouseholdId, &mExpense.Name, &mExpense.Amount)
	if err != nil {
		log.Fatal(err)
	}

	mExpense.Id = id

	return mExpense
}

func UpdateMonthlyExpense(mExpense models.MonthlyExpense) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("update monthly_expences set name = $1, amount = $2 where id = $3", mExpense.Name, mExpense.Amount, mExpense.Id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func DeleteMonthlyExpense(id int) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("delete from monthly_expences where id = $1", id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}
