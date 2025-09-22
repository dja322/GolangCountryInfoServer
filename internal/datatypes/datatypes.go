package datatypes

type AuthResult struct {
	ValidUser bool
	Calls     int
	CallLimit int
	Username  string
}

type ResponseType struct {
	ResponseData []byte
	ResponseCode int
}

type CountryDataType struct {
	GDP           int
	Population    int
	CapitolCity   string
	Continent     string
	SizeInSqMiles int
	Country       string
}
