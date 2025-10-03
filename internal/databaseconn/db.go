package databaseconn

/*
	Access database and get data
*/

import (
	"database/sql"
	"log"
)

var db *sql.DB //Database connection bool

var dataDrive string = "mysql"
var dataSoruce string = "admin:admin@tcp(127.0.0.1:3306)/"

func ConnectToDatabase() {
	var err error

	db, err = sql.Open(dataDrive, dataSoruce)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer db.Close()
}

func SelectFromDatabase(country string) (string, error) {

	var result string = ""
	var countryName string = "USA"
	//Queries
	//use db.Query for multiple rows
	err := db.QueryRow("Select * FROM country WHERE name=%s", countryName).Scan(&result)

	if err != nil {
		return "", err
	}

	log.Printf("Queryed %s, DATA: %s\n", country, result)

	return result, nil
}
