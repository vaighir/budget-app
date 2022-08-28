package db_helpers

import (
	"log"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
)

func CreateHousehold(household models.Household) int {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("insert into household (name) values ($1)", household.Name)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}

	var id int

	err = db.SQL.QueryRow("select currval(pg_get_serial_sequence('household','id'))").Scan(&id)
	if err != nil {
		panic(err)
	}

	return id

}

func GetHouseholdById(id int) models.Household {

	var household models.Household

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	err = db.SQL.QueryRow("select name, months_for_emergency_fund from household where id = $1", id).Scan(&household.Name, &household.MonthsOfEmergencyFund)
	if err != nil {
		panic(err)
	}

	household.Id = id

	return household
}

func UpdateHousehold(household models.Household) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer cleanup() 

	rows, err := db.SQL.Query("update household set months_for_emergency_fund = $1 where id = $2", household.MonthsOfEmergencyFund, household.Id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

	}
}
