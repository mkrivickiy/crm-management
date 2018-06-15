package method

type ShippingUpgrade struct {
	ShippingProfileID int     `json:"shipping_profile_id"` //Identifier for the parent shipping profile
	ValueID           int     `json:"value_id"`            //Identifier for the value
	Value             string  `json:"value"`               //Name of the shipping upgrade, e.g. USPS Priority
	Price             float32 `json:"price"`               //Additional cost of adding the shipping upgrade
	SecondaryPrice    float32 `json:"secondary_price"`     //Additional cost of adding the shipping upgrade with another item
	CurrencyCode      string  `json:"currency_code"`       //Currency for the price
	Type              int     `json:"type"`                //Domestic (0) or international (1)
	Order             int     `json:"order"`               //Display order
	Language          int     `json:"language"`            //Language code
}
