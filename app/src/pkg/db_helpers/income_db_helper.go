package db_helpers

import (
	"log"

	"github.com/vaighir/go-diet/app/pkg/drivers"
	"github.com/vaighir/go-diet/app/pkg/models"
)

func CreateIncome(householdId int, income models.Income) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("insert into income (household_id, name, amount) values ($1, $2, $3)", householdId, income.Name, income.Amount)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func GetAllIncomesByHouseholdId(householdId int) []models.Income {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("select id, name, amount from income where household_id = $1", householdId)
	if err != nil {
		log.Fatal(err)
	}

	var income models.Income
	var incomes []models.Income

	for rows.Next() {
		err := rows.Scan(&income.Id, &income.Name, &income.Amount)
		if err != nil {
			log.Println(err)
		}

		incomes = append(incomes, income)

	}

	return incomes
}

func GetIncomeById(id int) models.Income {

	var income models.Income

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select household_id, name, amount from income where id = $1", id).Scan(&income.HouseholdId, &income.Name, &income.Amount)
	if err != nil {
		log.Fatal(err)
	}

	income.Id = id

	return income
}

func DeleteIncome(id int) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("delete from income where id = $1", id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}
