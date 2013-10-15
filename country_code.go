package goiban

var (
	COUNTRY_TO_CC_MAP = map[string]string {
		"Albania": "AL",
		"Algeria": "DZ",
		"Andorra": "AD",
		"Angola": "AO",
		"Austria": "AT",
		"Azerbaijan": "AZ",
		"Bahrain": "BH",
		"Belgium": "BE",
		"Benin": "BJ",
		"Bosnia and Herzegovina": "BA",
		"Brazil": "BR",
		"British Virgin Islands": "VG",
		"Bulgaria": "BG",
		"Burkina Faso": "BF",
		"Burundi": "BI",
		"Cameroon": "CM",
		"Cape Verde": "CV",
		"Central African Republic": "FR",
		"Congo": "CG",
		"Costa Rica": "CR",
		"Croatia": "HR",
		"Cyprus": "CY",
		"Czech Republic": "CZ",
		"Denmark": "DK",
		"Dominican Republic": "DO",
		"Egypt": "EG",
		"Estonia": "EE",
		"Faroe Islands": "FO",
		"Finland": "FI",
		"France": "FR",
		"French Guiana": "FR",
		"French Polynesia": "FR",
		"Gabon": "GA",
		"Georgia": "GE",
		"Germany": "DE",
		"Gibraltar": "GI",
		"Greece": "GR",
		"Greenland": "GL",
		"Guadeloupe": "FR",
		"Guatemala": "GT",
		"Guernsey": "GB",
		"Hungary": "HU",
		"Iceland": "IS",
		"Iran": "IR",
		"Ireland": "IE",
		"Isle of Man": "GB",
		"Israel": "IL",
		"Italy": "IT",
		"Ivory Coast": "CI",
		"Jersey": "GB",
		"Kazakhstan": "KZ",
		"Kuwait": "KW",
		"Latvia": "LV",
		"Lebanon": "LB",
		"Liechtenstein": "LI",
		"Lithuania": "LT",
		"Luxembourg": "LU",
		"Macedonia": "MK",
		"Madagascar": "MG",
		"Mali": "ML",
		"Malta": "MT",
		"Martinique": "FR",
		"Mauritania": "MR",
		"Mauritius": "MU",
		"Moldova": "MD",
		"Monaco": "MC",
		"Montenegro": "ME",
		"Mozambique": "MZ",
		"Netherlands": "NL",
		"New Caledonia": "FR",
		"Norway": "NO",
		"Pakistan": "PK",
		"Palestine, State of": "PS",
		"Poland": "PL",
		"Portugal": "PT",
		"Romania": "RO",
		"Réunion": "FR",
		"Saint-Pierre and Miquelon": "FR",
		"San Marino": "SM",
		"Sao Tome and Principe": "PT",
		"Saudi Arabia": "SA",
		"Senegal": "SN",
		"Serbia": "RS",
		"Slovakia": "SK",
		"Slovenia": "SI",
		"Spain": "ES",
		"Sweden": "SE",
		"Switzerland": "CH",
		"Tunisia": "TN",
		"Turkey": "TR",
		"Ukraine": "UA",
		"United Arab Emirates": "AE",
		"United Kingdom": "GB",
		"Wallis and Futuna": "FR",
	}
)

/*
	Returns the allowed length of an IBAN code for a certain country.
			or -1 if the country code could not be looked up.
*/
func getCountryCode(countryName string) string {
	var countryCode string
	var ok bool

	countryCode, ok = COUNTRY_TO_CC_MAP[countryName]
	if ok {
		return countryCode
	} else {
		return ""
	}
}

