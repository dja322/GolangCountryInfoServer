package datatypes

type AuthResult struct {
	ValidUser bool
	Calls     int
	CallLimit int
	UserID    int
}

type AdminAuthResult struct {
	ValidAdmin bool
	AdminID    int
	AdminEmail string
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

type UserDataType struct {
	ID         int
	Tokenlimit int
	Tokenused  int
	Apikey     string
	Lastapiid  int
	Email      string
}

type DBConfig struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// Canonical country map: includes official names + common aliases and abbreviations
var CountryMap = map[string]string{
	// a
	"afghanistan": "afghanistan", "albania": "albania", "algeria": "algeria",
	"andorra": "andorra", "angola": "angola", "antigua and barbuda": "antigua and barbuda",
	"argentina": "argentina", "armenia": "armenia", "australia": "australia",
	"austria": "austria", "azerbaijan": "azerbaijan",

	// b
	"bahamas": "bahamas", "bahrain": "bahrain", "bangladesh": "bangladesh",
	"barbados": "barbados", "belarus": "belarus", "belgium": "belgium",
	"belize": "belize", "benin": "benin", "bhutan": "bhutan",
	"bolivia": "bolivia", "bosnia and herzegovina": "bosnia and herzegovina",
	"botswana": "botswana", "brazil": "brazil", "brunei": "brunei",
	"bulgaria": "bulgaria", "burkina faso": "burkina faso", "burundi": "burundi",

	// c
	"cabo verde": "cabo verde", "cape verde": "cabo verde",
	"cambodia": "cambodia", "cameroon": "cameroon", "canada": "canada",
	"central african republic": "central african republic", "chad": "chad",
	"chile": "chile", "china": "china", "colombia": "colombia",
	"comoros": "comoros",
	"congo":   "congo", "republic of the congo": "congo",
	"congo, republic of the":           "congo",
	"democratic republic of the congo": "drc", "drc": "drc",
	"costa rica": "costa rica", "croatia": "croatia", "cuba": "cuba",
	"cyprus": "cyprus", "czechia": "czechia", "czech republic": "czechia",

	// d–e
	"denmark": "denmark", "djibouti": "djibouti", "dominica": "dominica",
	"dominican republic": "dominican republic", "ecuador": "ecuador",
	"egypt": "egypt", "el salvador": "el salvador", "equatorial guinea": "equatorial guinea",
	"eritrea": "eritrea", "estonia": "estonia", "eswatini": "eswatini",
	"swaziland": "eswatini", "ethiopia": "ethiopia",

	// f–g
	"fiji": "fiji", "finland": "finland", "france": "france",
	"gabon": "gabon", "gambia": "gambia", "georgia": "georgia",
	"germany": "germany", "ghana": "ghana", "greece": "greece",
	"grenada": "grenada", "guatemala": "guatemala", "guinea": "guinea",
	"guinea-bissau": "guinea-bissau", "guyana": "guyana",

	// h–i
	"haiti": "haiti", "honduras": "honduras", "hungary": "hungary",
	"iceland": "iceland", "india": "india", "indonesia": "indonesia",
	"iran": "iran", "iraq": "iraq", "ireland": "ireland", "israel": "israel",
	"italy": "italy",

	// j–k
	"jamaica": "jamaica", "japan": "japan", "jordan": "jordan",
	"kazakhstan": "kazakhstan", "kenya": "kenya", "kiribati": "kiribati",
	"kuwait": "kuwait", "kyrgyzstan": "kyrgyzstan",

	// l
	"laos": "laos", "latvia": "latvia", "lebanon": "lebanon",
	"lesotho": "lesotho", "liberia": "liberia", "libya": "libya",
	"liechtenstein": "liechtenstein", "lithuania": "lithuania",
	"luxembourg": "luxembourg",

	// m
	"madagascar": "madagascar", "malawi": "malawi", "malaysia": "malaysia",
	"maldives": "maldives", "mali": "mali", "malta": "malta",
	"marshall islands": "marshall islands", "mauritania": "mauritania",
	"mauritius": "mauritius", "mexico": "mexico", "micronesia": "micronesia",
	"moldova": "moldova", "monaco": "monaco", "mongolia": "mongolia",
	"montenegro": "montenegro", "morocco": "morocco", "mozambique": "mozambique",
	"myanmar": "myanmar", "burma": "myanmar",

	// n
	"namibia": "namibia", "nauru": "nauru", "nepal": "nepal",
	"netherlands": "netherlands", "holland": "netherlands",
	"new zealand": "new zealand", "nicaragua": "nicaragua", "niger": "niger",
	"nigeria": "nigeria", "north korea": "north korea",
	"south korea": "south korea", "north macedonia": "north macedonia",
	"norway": "norway",

	// o–p
	"oman": "oman", "pakistan": "pakistan", "palau": "palau",
	"palestine": "palestine", "panama": "panama", "papua new guinea": "papua new guinea",
	"paraguay": "paraguay", "peru": "peru", "philippines": "philippines",
	"poland": "poland", "portugal": "portugal",

	// q–r
	"qatar": "qatar", "romania": "romania", "russia": "russia",
	"russian federation": "russia", "rwanda": "rwanda",

	// s
	"saint kitts and nevis": "saint kitts and nevis", "saint lucia": "saint lucia",
	"saint vincent and the grenadines": "saint vincent and the grenadines",
	"samoa":                            "samoa", "san marino": "san marino", "sao tome and principe": "sao tome and principe",
	"saudi arabia": "saudi arabia", "senegal": "senegal", "serbia": "serbia",
	"seychelles": "seychelles", "sierra leone": "sierra leone",
	"singapore": "singapore", "slovakia": "slovakia", "slovenia": "slovenia",
	"solomon islands": "solomon islands", "somalia": "somalia",
	"south africa": "south africa", "south sudan": "south sudan",
	"spain": "spain", "sri lanka": "sri lanka", "sudan": "sudan",
	"suriname": "suriname", "sweden": "sweden", "switzerland": "switzerland",
	"syria": "syria",

	// t
	"taiwan": "taiwan", "tajikistan": "tajikistan", "tanzania": "tanzania",
	"thailand": "thailand", "timor-leste": "timor-leste", "east timor": "timor-leste",
	"togo": "togo", "tonga": "tonga", "trinidad and tobago": "trinidad and tobago",
	"tunisia": "tunisia", "turkey": "turkey", "turkmenistan": "turkmenistan",
	"tuvalu": "tuvalu",

	// u
	"uganda": "uganda", "ukraine": "ukraine",
	"united arab emirates": "uae", "uae": "uae",
	"united kingdom": "uk", "uk": "uk", "england": "uk", "great britain": "uk",
	"united states": "usa", "united states of america": "usa", "us": "usa", "usa": "usa", "america": "usa",
	"uruguay": "uruguay", "uzbekistan": "uzbekistan",

	// v–z
	"vanuatu": "vanuatu", "vatican city": "vatican city", "holy see": "vatican city",
	"venezuela": "venezuela", "vietnam": "vietnam",
	"yemen": "yemen", "zambia": "zambia", "zimbabwe": "zimbabwe",
}
