package service

import (
	"GolangCountryInfoServer/internal/datatypes"
	"GolangCountryInfoServer/internal/server"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// https://go.dev/blog/json
func ParseRequest(query url.Values, userInfo datatypes.AuthResult) datatypes.ResponseType {

	//currently only parameter read is country
	var country string = query.Get("country")
	if !validCountry(country) {
		return datatypes.ResponseType{
			ResponseData: []byte("Error 400 Bad request, not a valid country"),
			ResponseCode: http.StatusBadRequest,
		}
	}

	response, err := server.GetCountryData(country)
	if err != nil {
		log.Printf("E: Error 500 internal server error %v", err)
		return datatypes.ResponseType{
			ResponseData: []byte("Error 500 internal server error"),
			ResponseCode: http.StatusInternalServerError,
		}
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Printf("E: Error 500 internal server error %v", err)
		return datatypes.ResponseType{
			ResponseData: []byte("Error 500 internal server error"),
			ResponseCode: http.StatusInternalServerError,
		}
	}

	return datatypes.ResponseType{
		ResponseData: []byte(jsonResponse),
		ResponseCode: http.StatusOK,
	}
}

func validCountry(country string) bool {
	_, exists := datatypes.CountryMap[strings.ToLower(strings.TrimSpace(country))]
	return exists
}
