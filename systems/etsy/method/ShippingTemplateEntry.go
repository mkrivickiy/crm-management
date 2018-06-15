package method

type ShippingTemplateEntry struct {
	ShippingTemplateEntryID int     `json:"shipping_template_entry_id"` //The numeric ID of this shipping template entry.
	ShippingTemplateID      int     `json:"shipping_template_id"`       //The numeric ID of the template this entry belongs to.
	CurrencyCode            string  `json:"currency_code"`              //The currency code for the prices in this template entry.
	OriginCountryID         int     `json:"origin_country_id"`          //The numeric ID of the country from which the listing ships.
	DestinationCountryID    int     `json:"destination_country_id"`     //The numeric ID of the country to which the listing ships (optional). If missing, these fees apply to all destinations.
	DestinationRegionID     int     `json:"destination_region_id"`      //The numeric ID of the region to which the listing ships (optional). Regions are collections of countries.
	PrimaryCost             float32 `json:"primary_cost"`               //The shipping fee for this item, if shipped alone.
	SecondaryCost           float32 `json:"secondary_cost"`             //The shipping fee for this item, if shipped with another item.
}
