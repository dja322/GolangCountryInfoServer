package server

import (
	DB_conn "GolangCountryInfoServer/internal/databaseconn"
	"GolangCountryInfoServer/internal/datatypes"
	"log"
)

/*
	Server functons

	Access user credentials in database
	Access database and pull country data

	process request and return data


*/

func GetCountryData(country string, specs ...string) datatypes.CountryDataType {
	//TODO: Access database and get data for specific country
	//TODO: call database functions to get data

	result, err := DB_conn.SelectFromDatabase(country)

	if err != nil {

	}

	log.Println(result)

	return datatypes.CountryDataType{
		GDP:           100000,
		Population:    1000000,
		CapitolCity:   "Test City",
		Continent:     "Test Continent",
		SizeInSqMiles: 10000,
		Country:       country,
	}
}

func GetUserData() string {
	return "data"
}
