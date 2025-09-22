package server

import "GolangCountryInfoServer/internal/datatypes"

/*
	Server functons

	Access user credentials in database
	Access database and pull country data

	process request and return data


*/

func GetCountryData(country string, specs ...string) datatypes.CountryDataType {
	//TODO: Access database and get data for specific country
	//TODO: call database functions to get data

	return datatypes.CountryDataType{}
}

func GetUserData() string {
	return "data"
}
