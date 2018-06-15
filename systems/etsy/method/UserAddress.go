package method

type UserAddress struct {
	UserAddressID     int    `json:"user_address_id"`     //The numeric ID of the user's address.
	UserID            int    `json:"user_id"`             //The user's numeric ID.
	Name              string `json:"name"`                //The user's name for this address.
	FirstLine         string `json:"first_line"`          //The first line of the user's address.
	SecondLine        string `json:"second_line"`         //The second line of the user's address.
	City              string `json:"city"`                //The city field of the user's address.
	State             string `json:"state"`               //The state field of the user's address.
	Zip               string `json:"zip"`                 //The zip code field of the user's address.
	CountryID         int    `json:"country_id"`          //The country's numeric ID.
	CountryName       string `json:"country_name"`        //The name of the user's country
	IsDefaultShipping bool   `json:"is_default_shipping"` //Is this the user's default shipping address
}
