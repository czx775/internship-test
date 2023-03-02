package models

type User struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"id" json:"name"`
	Age  string `form:"age" json:"age"`
}
