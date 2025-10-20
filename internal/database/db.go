package database

/*
	Access database and get data
*/

import (
	"GolangCountryInfoServer/internal/datatypes"
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil //Database connection var
var envFilePath string = "../.env"

func ConnectToDatabase() error {
	//if db not nil then db connection already established
	if db != nil {
		return nil
	}

	var err error

	//get environment variables from selected path
	var config datatypes.DBConfig = getENV(envFilePath)

	var cfg *mysql.Config = mysql.NewConfig()
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = config.Database

	var dataDrive string = "mysql"

	db, err = sql.Open(dataDrive, cfg.FormatDSN())
	if err != nil {
		log.Printf("ERROR: Error connecting to mysql: %v", err)
		return err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Printf("ERROR: Error connecting to the database: %v", err)
		return err
	}

	return nil
}

func SelectFromCountryDatabase(country string) (datatypes.CountryDataType, error) {

	err := ConnectToDatabase()
	if err != nil {
		return datatypes.CountryDataType{}, err
	}
	var id int = 0
	var data datatypes.CountryDataType

	//Queries
	//use db.Query for multiple rows
	// Use parameter placeholder (?) to avoid formatting issues and SQL injection
	err = db.QueryRow("SELECT * FROM Country WHERE name = ?", country).
		Scan(&id, &data.Country, &data.GDP, &data.Population, &data.CapitolCity, &data.Continent, &data.SizeInSqMiles)

	if err != nil {
		return datatypes.CountryDataType{}, err
	}
	log.Printf("INFO: Queryed %s, DATA ID: %d", country, id)

	return data, nil
}

func SelectFromUserDatabase(api_key string) (datatypes.UserDataType, error) {
	err := ConnectToDatabase()
	if err != nil {
		return datatypes.UserDataType{}, err
	}

	var data datatypes.UserDataType

	//Queries
	//use db.Query for multiple rows
	// Use parameter placeholder (?) to avoid formatting issues and SQL injection
	err = db.QueryRow("SELECT * FROM User WHERE apikey = ?", api_key).
		Scan(&data.ID, &data.Tokenlimit, &data.Tokenused, &data.Apikey, &data.Lastapiid, &data.Email)

	return data, err
}

func SelectFromAdminDatabase(email string, password string, passkey string) (datatypes.AdminAuthResult, error) {
	err := ConnectToDatabase()
	if err != nil {
		return datatypes.AdminAuthResult{}, err
	}

	var data datatypes.AdminAuthResult

	//Queries
	//use db.Query for multiple rows
	// Use parameter placeholder (?) to avoid formatting issues and SQL injection
	err = db.QueryRow("SELECT * FROM Admin WHERE email = ? AND password = ? AND passkey = ?",
		email, password, passkey).Scan(&data.AdminID, &data.AdminEmail)

	return data, err
}

/*
Admin functions are things that a user can request but never directly execute
*/
func Admin_SelectFromUserDatabase_generic(command string) (datatypes.UserDataType, error) {
	err := ConnectToDatabase()
	if err != nil {
		return datatypes.UserDataType{}, err
	}

	var data datatypes.UserDataType
	var commandFinal string = "SELECT * FROM User WHERE " + command

	err = db.QueryRow(commandFinal).
		Scan(&data.ID, &data.Tokenlimit, &data.Tokenused, &data.Apikey, &data.Lastapiid, &data.Email)

	return data, err
}

func Admin_AddUser(tokenlimit int, tokenused int, apikey string, lastapiid int, email string) (int, error) {
	err := ConnectToDatabase()
	if err != nil {
		return 0, err
	}

	//execute command on database
	result, err := db.Exec("INSERT INTO user (tokenlimit, tokenused, apikey, lastapiid, email) VALUES (?, ?, ?, ?, ?);",
		tokenlimit, tokenused, apikey, lastapiid, email)
	if err != nil {
		log.Printf("ERROR: Error adding user: %v", err)
	}

	var RowsAffected int64
	RowsAffected, _ = result.RowsAffected()
	return int(RowsAffected), err
}

func Admin_UpdateUser(id int, tokenlimit int, tokenused int, apikey string, lastapiid int, email string) (int, error) {
	err := ConnectToDatabase()
	if err != nil {
		return 0, err
	}

	//execute command on database
	result, err := db.Exec("UPDATE user SET tokenlimit=?, tokenused=?, apikey=?, lastapiid=?, email=? WHERE id=?;",
		tokenlimit, tokenused, apikey, lastapiid, email, id)
	if err != nil {
		log.Printf("ERROR: Error updating user: %v", err)
	}

	var RowsAffected int64
	RowsAffected, _ = result.RowsAffected()
	return int(RowsAffected), err
}

func Admin_RemoveUser(id int) (int, error) {
	err := ConnectToDatabase()
	if err != nil {
		return 0, err
	}

	//execute command on database
	result, err := db.Exec("DELETE FROM user WHERE id=?;", id)
	if err != nil {
		log.Printf("ERROR: Error removing user: %v", err)
	}

	var RowsAffected int64
	RowsAffected, _ = result.RowsAffected()
	return int(RowsAffected), err
}

// Initializing function for database, server should not start if this function fails
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

	log.Println("INFO: Initializing database with: ", cfg.FormatDSN()) //for credentials debug

	db, err = sql.Open(dataDrive, cfg.FormatDSN())
	if err != nil {
		log.Fatal("ERROR: Error connecting to mysql\n")
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("ERROR: Error connecting to the database: %v", err)
	}
	defer db.Close()

	//execute command on database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Database + ";")
	if err != nil {
		log.Fatalf("ERROR: Error creating database: %v", err)
	}

	var command string = "USE " + config.Database + ";"
	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("ERROR: Error Using the database: %v", err)
	}

	//Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS Country (
	id INT AUTO_INCREMENT PRIMARY KEY not NULL,
	name VARCHAR(255) not NULL,
	gdp INT not NULL,
	population INT not NULL,
	capitolcity VARCHAR (255) not NULL,
	continent VARCHAR (255) not NULL,
	sizeinsqmiles INT not NULL
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("ERROR: Error creating Country table %v\n", err)
	}

	//Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS User (
	id INT AUTO_INCREMENT PRIMARY KEY not NULL,
	tokenlimit INT not NULL,
	tokenused INT not NULL,
	apikey VARCHAR(255) not NULL,
	lastapiid INT not NULL,
	email VARCHAR(255) not NULL
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("ERROR: Error creating User table %v\n", err)
	}

	//Set up tables for database
	command = `CREATE TABLE IF NOT EXISTS Admin (
	id INT AUTO_INCREMENT PRIMARY KEY not NULL,
	password VARCHAR(255),
	passkey INT,
	email VARCHAR(255)
);`

	_, err = db.Exec(command)
	if err != nil {
		log.Fatalf("ERROR: Error creating Admin table %v\n", err)
	}

	return nil
}

//TODO Make function to verify database is set up properly when starting server but not remaking database
// , eg ensuring tables have correct fields and data is filled
//func VerifyDatabase

// Function to retrieve environment variables
func getENV(filepath string) datatypes.DBConfig {
	var config datatypes.DBConfig

	//get environment file
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("ERROR: Error opening file:", err)
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
		log.Println("ERROR: Error reading file:", err)
	}

	return config

}
