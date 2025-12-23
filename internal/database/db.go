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
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil //Database connection var
var envFilePath string = "../.env"

func ConnectToDatabase() error {
	//if db not nil then db connection already established
	if db != nil {
		return nil
	}

	//get environment variables from selected path
	var config datatypes.DBConfig = getENV(envFilePath)

	var cfg *mysql.Config = mysql.NewConfig()
	cfg.User = config.User
	cfg.Passwd = config.Password
	cfg.Net = "tcp"
	// Use host and port from config
	if config.Server == "" {
		config.Server = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = "3306"
	}
	cfg.Addr = config.Server + ":" + config.Port
	cfg.DBName = config.Database

	var dataDrive string = "mysql"

	// Try connecting with retries in case mysql is still initializing
	var err error
	const maxAttempts int = 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		db, err = sql.Open(dataDrive, cfg.FormatDSN())
		if err != nil {
			log.Printf("ERROR: Error opening DB (attempt %d/%d): %v", attempts+1, maxAttempts, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			return nil
		}

		log.Printf("WARN: DB ping failed (attempt %d/%d): %v", attempts+1, maxAttempts, err)
		_ = db.Close()
		time.Sleep(2 * time.Second)
	}

	log.Printf("ERROR: Unable to connect to DB after retries: %v", err)
	return err
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
	if config.Server == "" {
		config.Server = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = "3306"
	}
	cfg.Addr = config.Server + ":" + config.Port

	var dataDrive string = "mysql"

	log.Println("INFO: Initializing database with: ", cfg.FormatDSN()) //for credentials debug

	// Try opening DB with retries in case mysql is still initializing
	const maxAttempts = 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		db, err = sql.Open(dataDrive, cfg.FormatDSN())
		if err != nil {
			log.Printf("ERROR: Error opening DB for init (attempt %d/%d): %v", attempts+1, maxAttempts, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("WARN: DB ping failed during init (attempt %d/%d): %v", attempts+1, maxAttempts, err)
		_ = db.Close()
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("ERROR: Error connecting to the database after retries: %v", err)
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

	// If env file is loaded into container environment variables, use those
	if os.Getenv("DB_USER") != "" {
		config.User = os.Getenv("DB_USER")
		config.Password = os.Getenv("DB_PASSWORD")
		config.Database = os.Getenv("DB_NAME")
		config.Server = os.Getenv("DB_HOST")
		config.Port = os.Getenv("DB_PORT")

		if config.Server == "" {
			config.Server = "127.0.0.1"
		}
		if config.Port == "" {
			config.Port = "3306"
		}

		return config
	}

	// Fallback: read from .env file if present
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("ERROR: Error opening file:", err)
		return config
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "DB_SERVER=") {
			config.Server = strings.TrimPrefix(line, "DB_SERVER=")
		} else if strings.HasPrefix(line, "DB_PORT=") {
			config.Port = strings.TrimPrefix(line, "DB_PORT=")
		} else if strings.HasPrefix(line, "DB_USER=") {
			config.User = strings.TrimPrefix(line, "DB_USER=")
		} else if strings.HasPrefix(line, "DB_PASSWORD=") {
			config.Password = strings.TrimPrefix(line, "DB_PASSWORD=")
		} else if strings.HasPrefix(line, "DB_NAME=") {
			config.Database = strings.TrimPrefix(line, "DB_NAME=")
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Println("ERROR: Error reading file:", err)
	}

	// Defaults if values are missing
	if config.Server == "" {
		config.Server = "127.0.0.1"
	}
	if config.Port == "" {
		config.Port = "3306"
	}

	return config

}
