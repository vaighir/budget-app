package db_helpers

import (
	"log"
	"time"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func CreateUpcomingExpence(householdId int, uExpese models.UpcomingExpense) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("insert into upcoming_expences (household_id, name, amount, deadline) values ($1, $2, $3, $4)", householdId, uExpese.Name, uExpese.Amount, uExpese.Deadline)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}
}

func GetAllUpcomingExpensesByHouseholdId(householdId int) []models.UpcomingExpense {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("select id, name, amount, deadline from upcoming_expences where household_id = $1", householdId)
	if err != nil {
		panic(err)
	}

	var uExpese models.UpcomingExpense
	var uExpeses []models.UpcomingExpense

	for rows.Next() {
		err := rows.Scan(&uExpese.Id, &uExpese.Name, &uExpese.Amount, &uExpese.Deadline)
		if err != nil {
			panic(err)
		}

		uExpese.DeadlineString = uExpese.Deadline.Format("2006-01-02")

		uExpeses = append(uExpeses, uExpese)

	}

	return uExpeses
}

func GetUpcomingExpensesForHouseholdBetweenDates(householdId int, startDate time.Time, endDate time.Time) []models.UpcomingExpense {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("select id, name, amount, deadline from upcoming_expences where household_id = $1 and deadline between $2 and $3", householdId, startDate, endDate)
	if err != nil {
		panic(err)
	}

	var uExpese models.UpcomingExpense
	var uExpeses []models.UpcomingExpense

	for rows.Next() {
		err := rows.Scan(&uExpese.Id, &uExpese.Name, &uExpese.Amount, &uExpese.Deadline)
		if err != nil {
			panic(err)
		}

		uExpese.DeadlineString = uExpese.Deadline.Format("2006-01-02")

		uExpeses = append(uExpeses, uExpese)

	}

	return uExpeses
}

func GetUpcomingExpenseById(id int) models.UpcomingExpense {

	var uExpense models.UpcomingExpense

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	err = db.SQL.QueryRow("select household_id, name, amount, deadline from upcoming_expences where id = $1", id).Scan(&uExpense.HouseholdId, &uExpense.Name, &uExpense.Amount, &uExpense.Deadline)
	if err != nil {
		panic(err)
	}

	uExpense.DeadlineString = uExpense.Deadline.Format("2006-01-02")
	uExpense.Id = id

	return uExpense
}

func UpdateUpcomingExpense(uExpense models.UpcomingExpense) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("update upcoming_expences set name = $1, amount = $2, deadline = $3 where id = $4", uExpense.Name, uExpense.Amount, uExpense.Deadline, uExpense.Id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}
}

func DeleteUpcomingExpense(id int) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("delete from upcoming_expences where id = $1", id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}
}
