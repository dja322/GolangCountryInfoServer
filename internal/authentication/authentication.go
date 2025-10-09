package authentication

import (
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/server"
	"errors"
	"regexp"
)

/*
	Leave authetication for now, focus on getting base API functional

*/

// Checks if api-key is formatted correctly then checks if api key is present or valid
func AuthorizeUser(api_key string) (datatypes.AuthResult, error) {

	validKey, err := validateAPIKey(api_key)

	if err != nil {
		return datatypes.AuthResult{}, err
	}

	if !validKey {
		return datatypes.AuthResult{
			ValidUser: false,
			Calls:     -1,
			CallLimit: -1,
			UserID:    -1,
		}, nil
	}

	authResult, err := getUser(api_key)
	return authResult, err
}

// TODO: Implement authorization, this will call server.go function to actually access database credentials
// This will use hash and salt with SHA256 encryption eventually
func getUser(api_key string) (datatypes.AuthResult, error) {

	result, err := server.GetUserData(api_key)
	//test for internal server error or lack of data
	var ErrNoRows = errors.New("sql: no rows in result set")

	if err == ErrNoRows {
		var validUser bool = false
		return datatypes.AuthResult{
			ValidUser: validUser,
			Calls:     -1,
			CallLimit: -1,
			UserID:    -1,
		}, nil
	} else if err != nil {
		return datatypes.AuthResult{}, err
	}

	//For when user exists
	var validUser bool = true
	return datatypes.AuthResult{
		ValidUser: validUser,
		Calls:     result.Tokenused,
		CallLimit: result.Tokenlimit,
		UserID:    result.ID,
	}, nil
}

// this makes sure the API is not nefarious and is a proper key before accessing the database
func validateAPIKey(api_key string) (bool, error) {

	matched, err := regexp.MatchString("^gcs-\\d{5}-[A-Za-z0-9]{150,200}-[A-Za-z0-9]{4}$", api_key)

	//for temporary testing
	if api_key == "test" {
		matched = true
	}

	return matched, err

}
