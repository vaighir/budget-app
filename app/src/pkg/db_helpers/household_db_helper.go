package db_helpers

import (
	"log"

	"github.com/vaighir/go-diet/app/pkg/drivers"
	"github.com/vaighir/go-diet/app/pkg/models"
)

func CreateHousehold(household models.Household) int {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("insert into household (name) values ($1)", household.Name)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}

	var id int

	err = db.SQL.QueryRow("select currval(pg_get_serial_sequence('household','id'))").Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id

}
