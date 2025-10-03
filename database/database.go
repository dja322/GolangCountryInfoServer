package database

import (
	"database/sql"
	"log"
	"os"
)

type DBConfig struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

/*
TODO: This file should initialize and manage the database as separate from
server functions but in same repo for development
*/
var db *sql.DB //Database connection bool
var err error
var databaseName string = "country_db"

const logFileStr string = "../Logfile_database.log"

func InitializeDatabase() error {

	logFile, logErr := os.OpenFile(logFileStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatalf("error opening file: %v", logErr)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// config := &DBConfig{
	// 	Server:   "test",
	// 	Port:     "test",
	// 	User:     "test",
	// 	Password: "test",
	// 	Database: "test",
	// }

	// connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
	// 	config.Server,
	// 	config.Port,
	// 	config.User,
	// 	config.Password,
	// 	config.Database)

	log.Println("Initializing database")

	db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal("Error connecting to mysql\n")
	}
	defer db.Close()

	//execute command on database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)
	if err != nil {
		log.Fatal("Error creating database\n")
	}

	_, err = db.Exec("USE " + databaseName)
	if err != nil {
		log.Fatal("Error using database\n")
	}

	_, err = db.Exec("CREATE TABLE ")
	if err != nil {
		log.Fatal("Error creating table\n")
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
