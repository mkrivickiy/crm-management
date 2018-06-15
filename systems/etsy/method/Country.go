package method

type Country struct {
	CountryID            int     `json:"country_id"`              //The country's numeric ID.
	IsoCountryCode       string  `json:"iso_country_code"`        //The two-letter country code according to ISO 3166-1-alpha-2.
	WorldBankCountryCode string  `json:"world_bank_country_code"` //The three-letter country code according to the World Bank.
	Name                 string  `json:"name"`                    //The country's plain-English name.
	Slug                 string  `json:"slug"`                    //The country's plain-English name slugified; suitable for interpolation into a url.
	Lat                  float32 `json:"lat"`                     //The country's latitude.
	Lon                  float32 `json:"lon"`                     //The country's longitude.
}
