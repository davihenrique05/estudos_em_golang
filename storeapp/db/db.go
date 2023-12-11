package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectWithDb() *sql.DB {
	conectionString := "host=localhost user=root password=root dbname=storeapp port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conectionString)
	if err != nil {
		panic(err.Error())
	}

	return db
}
