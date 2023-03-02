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

	_, err := db.Exec("INSERT INTO user.user (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		log.Print(err)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users models.User
	var arrUsers []models.User

	db := databases.ConnectDatabase()
	defer db.Close()

	limit := 10
	offset := 0

	rows, err := db.Query("SELECT id, name, age FROM user.user LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&users.Id, &users.Name, &users.Age)
		if err != nil {
			log.Print(err)
		} else {
			arrUsers = append(arrUsers, users)
		}
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.FormValue("id")

	db := databases.ConnectDatabase()
	defer db.Close()

	row, err := db.Query("SELECT id, name, age FROM user.user WHERE id = ?", id)
	if err != nil {
		log.Print(err)
	}

	for row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Print(err)
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := databases.ConnectDatabase()
	defer db.Close()

	id := r.FormValue("id")
	name := r.FormValue("name")
	age := r.FormValue("age")

	_, err := db.Exec("UPDATE user.user SET name=?, age=? WHERE id=?", name, age, id)
	if err != nil {
		log.Print(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := databases.ConnectDatabase()
	defer db.Close()

	id := r.FormValue("id")

	_, err := db.Exec("DELETE FROM user.user WHERE id=?", id)
	if err != nil {
		log.Print(err)
	}

}

func GetUsersFiltered(w http.ResponseWriter, r *http.Request) {
	var users models.User
	var arrUsers []models.User

	db := databases.ConnectDatabase()
	defer db.Close()

	limit := 10
	offset := 0
	filter := r.FormValue("name")

	rows, err := db.Query("SELECT id, name, age FROM user.user WHERE name LIKE ? ORDER BY name LIMIT ? OFFSET ?", filter, limit, offset)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&users.Id, &users.Name, &users.Age)
		if err != nil {
			log.Print(err)
		} else {
			arrUsers = append(arrUsers, users)
		}
	}
}
