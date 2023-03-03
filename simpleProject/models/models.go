package models

type User struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"id" json:"name"`
	Age  string `form:"age" json:"age"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
	Row     User
}
