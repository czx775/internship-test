package controllers

import (
	"log"
	"net/http"
	"simpleProject/databases"
	"simpleProject/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := databases.ConnectDatabase()
	defer db.Close()

	name := r.FormValue("name")
	age := r.FormValue("age")

	_, err := db.Exec("INSERT INTO user(name, age) VALUES (?, ?)", name, age)
	if err != nil {
		log.Print(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users models.User
	var arrUsers []models.User

	db := databases.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age FROM user")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&users.ID, &users.Name, &users.Age)
		if err != nil {
			log.Print(err)
		} else {
			arrUsers = append(arrUsers, users)
		}
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func GetUsersFiltered(w http.ResponseWriter, r *http.Request) {

}
