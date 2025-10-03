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

type DBConfig struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}
