package database

/*
	Access database and get data
*/

import (
	"GolangCountryInfoServer/internal/datatypes"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB //Database connection bool
var envFilePath string = "../.env"

func ConnectToDatabase() {
	var err error

	var config datatypes.DBConfig = getENV(envFilePath) //Stub implement connection here soon

	var cfg *mysql.Config = mysql.NewConfig()
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

func AddUser() {

}

func AdminAddCountry() {

}

func AdminRemoveUser() {

}

func AdminAddAdmin() {

}

func InitializeDatabase() error {
	var db *sql.DB //Database connection bool
	var err error

	//get environment variables from selected path
	var config datatypes.DBConfig = getENV(envFilePath)

	var cfg *mysql.Config = mysql.NewConfig()
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	// cfg.DBName = config.Database

	var dataDrive string = "mysql"

	log.Println("Initializing database")
	// log.Println(cfg.FormatDSN()) //for credentials debug

	db, err = sql.Open(dataDrive, cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error connecting to mysql\n")
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	//execute command on database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Database + ";")
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	var command string = "USE " + config.Database + ";"
	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("Error Using the database: %v", err)
	}

	//TODO: Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS Country (
	id INT not NULL,
	name VARCHAR(255) PRIMARY KEY not NULL,
	gpd INT,
	population INT,
	capitolcity VARCHAR(255),
	sizeinsqmiles INT
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("Error creating Country table %v\n", err)
	}

	//TODO: Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS User (
	id INT PRIMARY KEY not NULL,
	tokenlimit INT not NULL,
	tokenused INT not NULL,
	apikey VARCHAR(255) not NULL,
	lastapiid INT,
	lastcall DATE,
	email VARCHAR(255)
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("Error creating User table %v\n", err)
	}

	//TODO: Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS Admin (
	id INT PRIMARY KEY not NULL,
	password VARCHAR(255),
	passkey INT,
	email VARCHAR(255)
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("Error creating Admin table %v\n", err)
	}

	log.Println("Successfully initialized database")

	return nil
}

func getENV(filepath string) datatypes.DBConfig {
	var config datatypes.DBConfig

	//get environment file
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Error opening file:", err)
		return config
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "DB_SERVER") {

			parts := strings.Split(scanner.Text(), "=")
			if len(parts) > 1 {
				config.Server = parts[1]
			} else {
				log.Fatal("Improperly formatted .env file, exiting.")
			}
		} else if strings.Contains(scanner.Text(), "DB_PORT") {
			parts := strings.Split(scanner.Text(), "=")
			if len(parts) > 1 {
				config.Port = parts[1]
			} else {
				log.Fatal("Improperly formatted .env file, exiting.")
			}
		} else if strings.Contains(scanner.Text(), "DB_USER") {
			parts := strings.Split(scanner.Text(), "=")
			if len(parts) > 1 {
				config.User = parts[1]
			} else {
				log.Fatal("Improperly formatted .env file, exiting.")
			}
		} else if strings.Contains(scanner.Text(), "DB_PASSWORD") {
			parts := strings.Split(scanner.Text(), "=")
			if len(parts) > 1 {
				config.Password = parts[1]
			} else {
				log.Fatal("Improperly formatted .env file, exiting.")
			}
		} else if strings.Contains(scanner.Text(), "DB_NAME") {
			parts := strings.Split(scanner.Text(), "=")
			if len(parts) > 1 {
				config.Database = parts[1]
			} else {
				log.Fatal("Improperly formatted .env file, exiting.")
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}

	return config

}
