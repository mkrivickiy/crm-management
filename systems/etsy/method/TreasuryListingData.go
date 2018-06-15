package method

type TreasuryListingData struct {
	UserID          int     `json:"user_id"`           //The numeric ID of the user who posted the item
	Title           string  `json:"title"`             //The listing's title
	Price           float32 `json:"price"`             //The item's price (private for sold listings)
	CurrencyCode    string  `json:"currency_code"`     //The ISO (alphabetic) code for the listing's currency
	ListingID       int     `json:"listing_id"`        //The ID of the listing
	State           string  `json:"state"`             //Whether the listing is active or not
	ShopName        string  `json:"shop_name"`         //The shop to which the listing belongs
	ImageID         int     `json:"image_id"`          //The ID of the main image of the listing
	ImageURL75x75   string  `json:"image_url_75x75"`   //The url to a 75x75 thumbnail of the main image.
	ImageURL170x135 string  `json:"image_url_170x135"` //The url to a 170x135 thumbnail of the main image.
}
