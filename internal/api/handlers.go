package api

import (
	"GolangCountryInfoServer/internal/authentication"
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/service"
	"log"
	"net/http"
)

/*
 * For general use
 * User will pass a country and the list of info they want
 */
func API_Base_Handler(w http.ResponseWriter, r *http.Request) {

	//sets initial write
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: Wrong HTTP Method used, use GET"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	//looks through request and if country key in address gets the value, if not return empty string
	var api_key string = query.Get("api_key")
	var authResult datatypes.AuthResult
	authResult, err := authentication.AuthorizeUser(api_key)

	//server errors
	if err != nil {
		log.Printf("ERROR: Error 500 internal server error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 internal server error"))
		return
	}

	//check if request is from valid user and has calls remaining
	if !authResult.ValidUser {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error 401 Unauthorized, Invalid API_key or unknown user"))
		log.Println("INFO: Unauthorized Access Attempt")
		return
	} else if authResult.CallLimit <= authResult.Calls {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Error 429 Too many requests, call limit reached for API key"))
		log.Println("INFO: Call limit reached by ", api_key)
		return
	}

	//Parse user request and return response object
	var response datatypes.ResponseType = service.ParseRequest(query, authResult)

	//pass response code and body back
	w.WriteHeader(response.ResponseCode)
	w.Write(response.ResponseData)

}

/*
 * This will handle admin functionality for modifying and fixing issues that don't require a hard reset
 * of the server or alterations to code
 */
func Admin_Handler(w http.ResponseWriter, r *http.Request) {
	//sets initial write
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: Wrong HTTP Method used, use POST"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	//looks through request and if country key in address gets the value, if not return empty string
	var password string = query.Get("password")
	var passkey string = query.Get("passkey")
	var email string = query.Get("email")
	var authResult datatypes.AdminAuthResult
	authResult, err := authentication.AuthorizeAdmin(password, passkey, email)

	//server errors
	if err != nil {
		log.Printf("ERROR: Error 500 internal server error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 internal server error"))
		return
	}

	//check if request is from valid admin
	if !authResult.ValidAdmin {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error 401 Unauthorized, Invalid password/passkey or unknown admin"))
		log.Println("INFO: Unauthorized Admin Access Attempt")
		return
	}

	//Parse user request and return response object
	var response datatypes.ResponseType = service.ParseAdminRequest(query, authResult)

	//pass response code and body back
	w.WriteHeader(response.ResponseCode)
	w.Write(response.ResponseData)
}

// Serves the root landing page
func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/index.html")
}
