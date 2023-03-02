package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:nicolas1@tcp(127.0.0.1:3306)/simpleProject")
	if err != nil {
		panic(err.Error())
	}
	return db
}
