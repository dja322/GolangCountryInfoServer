package main

/*
TODOLIST:

Add additional request parameters for users, whenever priority
Add docker containerization
Add hashing/salting
	//encrypting salting for user security
	//https://en.wikipedia.org/wiki/Salt_%28cryptography%29#Since_the_1980s
	// TODO: Implement authorization hash and salt with SHA256 encryption
	// eventually when user is registering

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
const port string = ":3000"

func main() {
	logFile, logErr := os.OpenFile(logFileStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Fatalf("ERROR: error opening file: %v", logErr)
	}
	defer logFile.Close()

	//sets global output log file for whole project
	log.SetOutput(logFile)

	log.Println("INFO: Log file Created")

	//initialize database if set
	if initDatabase {
		err := database.InitializeDatabase()
		if err != nil {
			log.Fatalf("error initializing database: %v", err)
		}
		log.Println("INFO: Successfully initialzed database")
	}

	//began setting up server
	fmt.Println("INFO: Starting Server...")

	setHandlers() //sets the handlers for different endpoints

	log.Println("INFO: API handlers set")

	serverErr := http.ListenAndServe(port, nil) // Start the server
	if serverErr != nil {
		fmt.Println("Error starting server:", serverErr)
		log.Printf("ERROR: Error starting server %v", serverErr)
	}
	log.Println("INFO: Server terminated")

}

/*
 * Sets API handlers for routing requests
 */
func setHandlers() {

	http.HandleFunc("/api/", api.API_Base_Handler)
	http.HandleFunc("/admin/", api.Admin_Handler)
	http.HandleFunc("/", api.RootHandler)
}
