package method

type Variations_Option struct {
	ValueID        int    `json:"value_id"`        //The numeric ID of the option
	Value          string `json:"value"`           //The normalized value of the option
	FormattedValue string `json:"formatted_value"` //The formatted/translated value of the option
	IsAvailable    bool   `json:"is_available"`    //True if the option is available for purchase
	PriceDiff      int    `json:"price_diff"`      //The additional price, if any, for this option above the Listing's base price. (NOTE: Unavailable options may have negative values.)
	Price          int    `json:"price"`           //The full price of this option, added to the Listing's base price. (NOTE: Unavailable options may have a price that is lower than the Listing's base price, which is based on the lowest available option price.)
}
