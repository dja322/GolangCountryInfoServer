package authentication

import (
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/server"
)

/*
	Leave authetication for now, focus on getting base API functional

*/

func AuthorizeUser(api_key string) datatypes.AuthResult {

	if !validateAPIKey(api_key) {
		return datatypes.AuthResult{
			ValidUser: false,
			Calls:     -1,
			CallLimit: -1,
			Username:  "Invalid",
		}
	}

	return getUser(api_key)
}

// TODO: Implement authorization, this will call server.go function to actually access database credentials
// This will use hash and salt with SHA256 encryption
func getUser(api_key string) datatypes.AuthResult {

	//TODO: Implement user data and getting user
	var validUser bool = true
	var calls int = 1
	var callLimit int = 100
	var user string = "nil"

	if api_key == "" {
		validUser = false
	} else {
		user = server.GetUserData()
	}

	//for server import delete later

	return datatypes.AuthResult{
		ValidUser: validUser,
		Calls:     calls,
		CallLimit: callLimit,
		Username:  user,
	}
}

// this makes sure the API is not nefarious and is a proper key before accessing the database
func validateAPIKey(api_key string) bool {

	if api_key == "evil" {
		print("")
		return false
	}
	return true
}
