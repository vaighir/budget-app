package db_helpers

import (
	"log"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func CreateFund(householdId int, fund models.Fund) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("insert into funds (household_id, name, amount) values ($1, $2, $3)", householdId, fund.Name, fund.Amount)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func GetAllFundsByHouseholdId(householdId int) []models.Fund {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	
	defer closeConnectionAndRecover() 

	rows, err := db.SQL.Query("select id, name, amount from funds where household_id = $1", householdId)
	if err != nil {
		log.Fatal(err)
	}

	var fund models.Fund
	var funds []models.Fund

	for rows.Next() {
		err := rows.Scan(&fund.Id, &fund.Name, &fund.Amount)
		if err != nil {
			panic(err)
		}

		funds = append(funds, fund)

	}

	return funds
}

func GetFundById(id int) models.Fund {

	var fund models.Fund

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select household_id, name, amount from funds where id = $1", id).Scan(&fund.HouseholdId, &fund.Name, &fund.Amount)
	if err != nil {
		log.Fatal(err)
	}

	fund.Id = id

	return fund
}

func UpdateFund(fund models.Fund) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("update funds set name = $1, amount = $2 where id = $3", fund.Name, fund.Amount, fund.Id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func DeleteFund(id int) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("delete from funds where id = $1", id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func closeConnectionAndRecover(){
	db.SQL.Close()

		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
		}
}