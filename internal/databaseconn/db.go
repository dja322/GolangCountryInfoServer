package databaseconn

/*
	Access database and get data
*/

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

var db *sql.DB //Database connection bool

func ConnectToDatabase() {
	var err error

	var config DBConfig // = getENV(envFilePath) Stub implement connection here soon

	cfg := mysql.NewConfig()
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = config.Database

	var dataDrive string = "mysql"
	var dataSource string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open(dataDrive, dataSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
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
