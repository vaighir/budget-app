package db_helpers

import (
	"log"

	"github.com/vaighir/go-diet/app/pkg/drivers"
	"github.com/vaighir/go-diet/app/pkg/models"
)

const dbDns = "host=localhost port=5432 dbname=test_db user=user password=password"

func GetAllUsers() []models.User {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("select id, username, password from users")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var username, password string
	var user models.User

	var users []models.User

	for rows.Next() {
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			log.Println(err)
		}

		user.Id = id
		user.Username = username
		user.Password = password

		users = append(users, user)

	}

	return users

}

func GetUserById(id int) models.User {

	var username, password string
	var user models.User

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select username, password from users where id = $1", id).Scan(&username, &password)
	if err != nil {
		log.Println(err)
	}

	user.Id = id
	user.Username = username
	user.Password = password

	return user
}

func UpdateUser(models.User) {

}

func CreateUser(models.User) {

}

func DeleteUser(models.User) {

}
