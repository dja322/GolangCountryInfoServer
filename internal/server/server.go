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
	result, err := database.SelectFromDatabase(country)

	if err != nil {
		return datatypes.CountryDataType{}, err
	}

	return datatypes.CountryDataType{
		GDP:           result.GDP,
		Population:    result.Population,
		CapitolCity:   result.CapitolCity,
		Continent:     result.Continent,
		SizeInSqMiles: result.SizeInSqMiles,
		Country:       result.Country,
	}, nil
}

func GetUserData() string {
	return "data"
}
