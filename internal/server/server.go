package server

import (
	"GolangCountryInfoServer/internal/database"
	"GolangCountryInfoServer/internal/datatypes"
	"log"
)

/*
	Server functons

	Access user credentials in database
	Access database and pull country data

	process request and return data


*/

func GetCountryData(country string, specs ...string) (datatypes.CountryDataType, error) {
	//TODO: Access database and get data for specific country
	//TODO: call database functions to get data

	//get country data from database if it exists
	log.Printf("S: Querying For country: %s", country)
	result, err := database.SelectFromCountryDatabase(country)

	return result, err
}

func GetUserData(apikey string) (datatypes.UserDataType, error) {

	//get country data from database if it exists
	log.Printf("S: Getting data for key: %s", apikey)
	result, err := database.SelectFromUserDatabase(apikey)

	return result, err
}
