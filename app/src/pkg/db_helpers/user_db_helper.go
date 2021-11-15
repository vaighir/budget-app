package db_helpers

import (
	"log"

	"github.com/vaighir/budget-app/app/pkg/drivers"
	"github.com/vaighir/budget-app/app/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() []models.User {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("select id, username, password, household_id from users")
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	var users []models.User

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.HouseholdId)
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

	err = db.SQL.QueryRow("select username, password, household_id from users where id = $1", id).Scan(&user.Username, &user.Password, &user.HouseholdId)
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

	err = db.SQL.QueryRow("select id, password, household_id from users where username = $1", username).Scan(&user.Id, &user.Password, &user.HouseholdId)
	if err != nil {
		log.Println(err)
	}

	user.Username = username

	return user
}

func UpdateUser(models.User) {

}

func AddHouseholdToUser(user models.User, householdId int) {
	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	rows, err := db.SQL.Query("update users set household_id = $1 where id = $2", householdId, user.Id)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}
}

func RemoveHouseholdFromUser(u models.User, householdId int) {

}

func CreateUser(user models.User) {

	db, err := drivers.ConnectSQL(dbDns)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
		return
	}

	rows, err := db.SQL.Query("insert into users (username, password) values ($1, $2)", user.Username, hashedPassword)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {

	}

}

func DeleteUser(models.User) {

}
