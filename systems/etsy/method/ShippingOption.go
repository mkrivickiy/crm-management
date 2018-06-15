package method

type ShippingOption struct {
	OptionID     string `json:"option_id"`     //The ID of the shipping option.
	Name         string `json:"name"`          //The name of the shipping option.
	OptionType   int    `json:"option_type"`   //The type of shipping option.
	Cost         string `json:"cost"`          //The total cart shipping price if the shipping option is selected.
	CurrencyCode string `json:"currency_code"` //The ISO (alphabetic) code for the shipping option's currency.
}
