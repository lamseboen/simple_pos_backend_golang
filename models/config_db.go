package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	host := "localhost"
	user := "simple_pos"
	password := "simple_pos"
	dbname := "db_simple_pos"
	port := 5118

	pgCon := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		port, host, user, password, dbname)

	db, err := sql.Open("postgres", pgCon)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")
	return db
}
