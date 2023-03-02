package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"simpleProject/controllers"
)

func main() {
	router := mux.NewRouter()

	// CRUD endpoints
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Filtering and pagination endpoint
	router.HandleFunc("/users/filtered", controllers.GetUsersFiltered).Methods("GET")

	http.ListenAndServe(":8000", router)
}
