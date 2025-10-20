package server

import (
	"GolangCountryInfoServer/internal/database"
	"GolangCountryInfoServer/internal/datatypes"
	"log"
	"net/url"
	"strconv"
)

/*
	Server functons

	Access user credentials in database
	Access database and pull country data

	process request and return data
*/

func GetCountryData(country string, specs ...string) (datatypes.CountryDataType, error) {

	//get country data from database if it exists
	log.Printf("INFO: Querying For country: %s", country)
	result, err := database.SelectFromCountryDatabase(country)

	return result, err
}

func GetUserData(apikey string) (datatypes.UserDataType, error) {

	//get user data from database if it exists
	log.Printf("INFO: Getting data for key: %s", apikey)
	result, err := database.SelectFromUserDatabase(apikey)

	return result, err
}

func GetAdminData(email string, password string, passkey string) (datatypes.AdminAuthResult, error) {

	//get admin data from database if it exists
	log.Printf("INFO: Getting admin data for email: %s", email)
	result, err := database.SelectFromAdminDatabase(email, password, passkey)

	return result, err
}

func Admin_getUserData_generic(command string) (datatypes.UserDataType, error) {

	//get user data from database if it exists
	log.Printf("INFO: Admin querying user data with command: %s", command)
	result, err := database.Admin_SelectFromUserDatabase_generic(command)

	return result, err
}

func ResolveAdminRequest(query url.Values, purpose string, adminInfo datatypes.AdminAuthResult) (int, error) {
	switch purpose {
	case "add_user":
		//add user functionality
		tokenlimit, _ := strconv.Atoi(query.Get("tokenlimit"))
		tokenused, _ := strconv.Atoi(query.Get("tokenused"))
		lastapiid, _ := strconv.Atoi(query.Get("lastapiid"))
		result, err := database.Admin_AddUser(tokenlimit, tokenused, query.Get("apikey"),
			lastapiid, query.Get("email"))
		return result, err

	case "remove_user":
		//remove user functionality
		userID, _ := strconv.Atoi(query.Get("userID"))
		result, err := database.Admin_RemoveUser(userID)
		return result, err
	case "update_user":
		//update user functionality
		id, _ := strconv.Atoi(query.Get("id"))
		tokenlimit, _ := strconv.Atoi(query.Get("tokenlimit"))
		tokenused, _ := strconv.Atoi(query.Get("tokenused"))
		lastapiid, _ := strconv.Atoi(query.Get("lastapiid"))
		result, err := database.Admin_UpdateUser(id, tokenlimit, tokenused, query.Get("apikey"),
			lastapiid, query.Get("email"))
		return result, err
	}

	return 0, nil
}
