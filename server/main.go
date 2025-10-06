package main

//encrypting salting for user security
//https://en.wikipedia.org/wiki/Salt_%28cryptography%29#Since_the_1980s

/*
TODOLIST:

Implement database access and functionality
Implement base api request through server, high priority
Add additional request parameters for users, whenever priority
Implement authetication, low priority do after base api and database functionality works

*/

import (
	"GolangCountryInfoServer/internal/api"
	"GolangCountryInfoServer/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"
)

const logFileStr string = "../Logfile_server.log"
const initDatabase bool = true

func main() {
	logFile, logErr := os.OpenFile(logFileStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatalf("error opening file: %v", logErr)
	}
	defer logFile.Close()

	//sets global output log file for whole project
	log.SetOutput(logFile)

	log.Println("Log file Created")

	//initialize database if set
	if initDatabase {
		err := database.InitializeDatabase()
		if err != nil {
			log.Fatalf("error initializing database: %v", err)
		}
		log.Println("Successfully initialzed database")
	}

	//began setting up server
	fmt.Println("Starting Server...")

	//set up server logfile

	setHandlers() //sets the handlers for different endpoints

	log.Println("API handlers set")

	log.Println("Server started")
	serverErr := http.ListenAndServe(":3000", nil) // Start the server
	if serverErr != nil {
		fmt.Println("Error starting server:", serverErr)
	}
	log.Println("Server terminated")

}

/*
 * Sets API handlers for routing requests
 */
func setHandlers() {

	http.HandleFunc("/api/", api.API_Base_Handler)
	http.HandleFunc("/admin", api.Admin_Handler)
}
