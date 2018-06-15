package method

type ShippingInfo struct {
	ShippingInfoID         int     `json:"shipping_info_id"`         //The numeric ID of this shipping info record.
	OriginCountryID        int     `json:"origin_country_id"`        //The numeric ID of the country from which the listing ships.
	DestinationCountryID   int     `json:"destination_country_id"`   //The numeric ID of the country to which the listing ships (optional). If missing, these fees apply to all destinations.
	CurrencyCode           string  `json:"currency_code"`            //Identifies the currency unit for shipping fees (currently, always 'USD').
	PrimaryCost            float32 `json:"primary_cost"`             //The shipping fee for this item, if shipped alone.
	SecondaryCost          float32 `json:"secondary_cost"`           //The shipping fee for this item, if shipped with another item.
	ListingID              int     `json:"listing_id"`               //The numeric ID of the listing to which this shipping info applies.
	RegionID               int     `json:"region_id"`                //The numeric ID of the region to which this shipping info applies (optional). If missing, no region is selected. Regions are shorthand for lists of individual countries.
	OriginCountryName      string  `json:"origin_country_name"`      //The name of the country from which this item ships.
	DestinationCountryName string  `json:"destination_country_name"` //The name of the country to which this item ships.
}
