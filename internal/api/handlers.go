package api

import (
	"GolangCountryInfoServer/internal/authentication"
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/service"
	"fmt"
	"log"
	"net/http"
)

/*
 * For general use
 * User will pass a country and the list of info they want
 */
func API_Base_Handler(w http.ResponseWriter, r *http.Request) {

	//sets initial w rite
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	//looks through request and if country key in address gets the value, if not return empty string
	//example of how requests look on other api https://apidocs.cheapshark.com/#b9b738bf-2916-2a13-e40d-d05bccdce2ba
	var api_key string = query.Get("api_key")
	var authResult datatypes.AuthResult = authentication.AuthorizeUser(api_key)

	//check if request is from valid user and has calls remaining
	if !authResult.ValidUser {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error 401 Unauthorized"))
		log.Println("Unauthorized Access Attempt")
		return
	} else if authResult.CallLimit <= authResult.Calls {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Error 429 Too many requests"))
		log.Println("Call limit reached by ", api_key)
		return
	}

	//Parse user request and return response object
	log.Println("Request made by user", authResult.Username)
	var response datatypes.ResponseType = service.ParseRequest(r, authResult)
	log.Println("User request processed for user", authResult.Username)

	//pass response code and body back
	w.WriteHeader(response.ResponseCode)
	w.Write(response.ResponseData)

}

/*
 * This will handle admin functionality for modifying and fixing issues that don't require a hard reset
 * of the server or alterations to code
 */
func Admin_Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Admin!")
}
