package models

import "database/sql"

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.99.100:3306)/go_db")

	if err != nil {
		panic(err.Error())
	}

	return db
}
