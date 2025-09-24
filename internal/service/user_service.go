package service

import (
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/server"
	"encoding/json"
	"log"
	"net/http"
)

// https://go.dev/blog/json
func ParseRequest(r *http.Request, userInfo datatypes.AuthResult) datatypes.ResponseType {
	//TODO: Parse requests for user that will then be sent to server.go which will actually access the database
	query := r.URL.Query()

	//currently only parameter read is country
	var country string = query.Get("country")
	if !validCountry(country) {
		return datatypes.ResponseType{
			ResponseData: []byte("Error 400 Bad request"),
			ResponseCode: http.StatusBadRequest,
		}
	}

	log.Println("Queryied for data from ", country)
	var response datatypes.CountryDataType = server.GetCountryData(country)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return datatypes.ResponseType{
			ResponseData: []byte("Error 500 internal server error"),
			ResponseCode: http.StatusInternalServerError,
		}
	}

	log.Println("Successful request for", country)
	return datatypes.ResponseType{
		ResponseData: []byte(jsonResponse),
		ResponseCode: http.StatusOK,
	}
}

// keys := make([]string, 0, len(query))
// for k := range query {
// 	keys = append(keys, k)
// }

// for key, values := range query {
// 	for _, value := range values {
// 		fmt.Print(w, "Key: %s, Value: %s\n", key, value)
// 	}
// }

func validCountry(country string) bool {
	if country == "y" {
		print("yes")
		return false
	}
	return true
}
