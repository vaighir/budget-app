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

	var user models.User
	var users []models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			log.Println(err)
		}

		users = append(users, user)

	}

	return users

}

func GetUserById(id int) models.User {

	var user models.User

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select username, password from users where id = $1", id).Scan(&user.Username, &user.Password)
	if err != nil {
		log.Println(err)
	}

	user.Id = id

	return user
}

func GetUserByUsername(username string) models.User {

	var user models.User

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	err = db.SQL.QueryRow("select id, password from users where username = $1", username).Scan(&user.Id, &user.Password)
	if err != nil {
		log.Println(err)
	}

	user.Username = username

	return user
}

func UpdateUser(models.User) {

}

func CreateUser(models.User) {

}

func DeleteUser(models.User) {

}
