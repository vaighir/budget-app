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

//func GetAllIncomesByHouseholdId(householdId int) []models.Income {}
