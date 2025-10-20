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

// Partse user request and send proper request to server
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
		log.Printf("ERROR: Error 500 internal server error %v", err)
		return datatypes.ResponseType{
			ResponseData: []byte("Error 500 internal server error"),
			ResponseCode: http.StatusInternalServerError,
		}
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Printf("ERROR: Error 500 internal server error %v", err)
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

func ParseAdminRequest(query url.Values, adminInfo datatypes.AdminAuthResult) datatypes.ResponseType {
	purpose := query.Get("purpose")

	if purpose == "get_user_data_generic" {
		response, err := server.Admin_getUserData_generic(query.Get("command"))
		if err != nil {
			log.Printf("ERROR: Error 500 internal server error %v", err)
			return datatypes.ResponseType{
				ResponseData: []byte("Error 500 internal server error"),
				ResponseCode: http.StatusInternalServerError,
			}
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Printf("ERROR: Error 500 internal server error %v", err)
			return datatypes.ResponseType{
				ResponseData: []byte("Error 500 internal server error"),
				ResponseCode: http.StatusInternalServerError,
			}
		}
		return datatypes.ResponseType{
			ResponseData: []byte(jsonResponse),
			ResponseCode: http.StatusOK,
		}
	} else {
		rowsEffected, err := server.ResolveAdminRequest(query, purpose, adminInfo)
		if err != nil {
			log.Printf("ERROR: Error 500 internal server error %v", err)
			return datatypes.ResponseType{
				ResponseData: []byte("Error 500 internal server error"),
				ResponseCode: http.StatusInternalServerError,
			}
		} else {
			return datatypes.ResponseType{
				ResponseData: []byte("Success" + string(rune(rowsEffected)) + " rows affected"),
				ResponseCode: http.StatusOK,
			}
		}
	}

}
