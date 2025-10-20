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
		tokenlimit, err := strconv.Atoi(query.Get("tokenlimit"))
		if err != nil {
			return 0, err
		}
		tokenused, err := strconv.Atoi(query.Get("tokenused"))
		if err != nil {
			return 0, err
		}
		lastapiid, err := strconv.Atoi(query.Get("lastapiid"))
		if err != nil {
			return 0, err
		}
		apikey := query.Get("apikey")
		email := query.Get("email")
		result, err := database.Admin_AddUser(tokenlimit, tokenused, apikey,
			lastapiid, email)
		return result, err

	case "remove_user":
		//remove user functionality
		userID, err := strconv.Atoi(query.Get("userID"))
		if err != nil {
			return 0, err
		}
		result, err := database.Admin_RemoveUser(userID)
		return result, err
	case "update_user":
		//update user functionality
		id, err := strconv.Atoi(query.Get("id"))
		if err != nil {
			return 0, err
		}
		tokenlimit, err := strconv.Atoi(query.Get("tokenlimit"))
		if err != nil {
			return 0, err
		}
		tokenused, err := strconv.Atoi(query.Get("tokenused"))
		if err != nil {
			return 0, err
		}
		lastapiid, err := strconv.Atoi(query.Get("lastapiid"))
		if err != nil {
			return 0, err
		}
		apikey := query.Get("apikey")
		email := query.Get("email")
		result, err := database.Admin_UpdateUser(id, tokenlimit, tokenused, apikey,
			lastapiid, email)
		return result, err
	}

	return 0, nil
}
