package method

type UserProfile struct {
	UserProfileID        int     `json:"user_profile_id"`        //The numeric ID of the user profile. Deprecated: Do not use
	UserID               int     `json:"user_id"`                //The numeric ID of the user this profile corresponds to.
	LoginName            string  `json:"login_name"`             //The login name of the user that owns this profile.
	Bio                  string  `json:"bio"`                    //The bio of the user.
	Gender               string  `json:"gender"`                 //The gender of the user.
	BirthMonth           string  `json:"birth_month"`            //The birth month of the user.
	BirthDay             string  `json:"birth_day"`              //The birth day of the user.
	BirthYear            string  `json:"birth_year"`             //The birth year of the user.
	JoinTsz              float32 `json:"join_tsz"`               //The date and time the user was created, in epoch seconds.
	Materials            string  `json:"materials"`              //A list of the user's favorite materials.
	CountryID            int     `json:"country_id"`             //The numeric ID of the user's country.
	Region               string  `json:"region"`                 //The user's region (free-form text)
	City                 string  `json:"city"`                   //The user's city (free-form text)
	Location             string  `json:"location"`               //The user's location (self-supplied, free-form text). Deprecated: Do not use.
	AvatarID             int     `json:"avatar_id"`              //The numeric ID of this user's avatar image.
	Lat                  float32 `json:"lat"`                    //The latitude of the user.
	Lon                  float32 `json:"lon"`                    //The longitude of the user.
	TransactionBuyCount  int     `json:"transaction_buy_count"`  //The number of bought items for this user.
	TransactionSoldCount int     `json:"transaction_sold_count"` //The number of sold transactions for this user.
	IsSeller             bool    `json:"is_seller"`              //If the user is a seller or not. Value true if user is a seller.
	ImageURL75x75        string  `json:"image_url_75x75"`        //The URL to the user's avatar image; dimensions 75px square.
	FirstNname           string  `json:"first_name"`             //The first name of the user.
	LastName             string  `json:"last_name"`              //The last name of the user.
}
