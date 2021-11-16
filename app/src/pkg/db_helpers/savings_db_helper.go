package db_helpers

import (
	"log"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func CreateSavings(householdId int, savings models.Savings) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("insert into savings (household_id, name, amount) values ($1, $2, $3)", householdId, savings.Name, savings.Amount)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func GetAllSavingsByHouseholdId(householdId int) []models.Savings {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("select id, name, amount from savings where household_id = $1", householdId)
	if err != nil {
		log.Fatal(err)
	}

	var savings models.Savings
	var savingsList []models.Savings

	for rows.Next() {
		err := rows.Scan(&savings.Id, &savings.Name, &savings.Amount)
		if err != nil {
			log.Println(err)
		}

		savingsList = append(savingsList, savings)

	}

	return savingsList
}

func GetSavingsById(id int) models.Savings {

	var savings models.Savings

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select household_id, name, amount from savings where id = $1", id).Scan(&savings.HouseholdId, &savings.Name, &savings.Amount)
	if err != nil {
		log.Fatal(err)
	}

	savings.Id = id

	return savings
}

func UpdateSavings(savings models.Savings) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("update savings set name = $1, amount = $2 where id = $3", savings.Name, savings.Amount, savings.Id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func DeleteSavings(id int) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("delete from savings where id = $1", id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}
