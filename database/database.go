package database

import (
	"database/sql"
)

var db *sql.DB //Database connection bool

func InitializeDatabase() error {

	var err error

	db, err = sql.Open("mysql", "")

	if err != nil {

	}
	defer db.Close()

	//executes command on DB
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS COUNTRY_DB")
	if err != nil {

	}

	var result string = ""
	var countryName string = "USA"
	//Queries
	//use db.Query for multiple rows
	err = db.QueryRow("Select * FROM country WHERE name=%s", countryName).Scan(&result)
	if err != nil {

	}

	return nil
}
