package method

type CartListing struct {
	ListingID              int                           `json:"listing_id"`               //The numeric ID of the listing
	PurchaseQuantity       int                           `json:"purchase_quantity"`        //The quantity of the listing being purchased
	PurchaseState          string                        `json:"purchase_state"`           //The purchase state of the listing, possible values: valid, invalid_quantity, invalid_shipping, not_active, edited, invalid_currency, invalid_shipping_currency
	IsDigital              bool                          `json:"is_digital"`               //True if this listing is for a digital download.
	FileData               string                        `json:"file_data"`                //Written description of the files attached to this digital listing.
	ListingCustomizationID int                           `json:"listing_customization_id"` //When Variations are present on a listing, this can be used to differentiate multiple instances of the same listing.
	VariationsAvailable    bool                          `json:"variations_available"`     //Whether Variations are available for this listing.
	HasVariations          bool                          `json:"has_variations"`           //Whether the buyer selected Variations for this listing.
	SelectedVariations     []Variations_SelectedProperty `json:"selected_variations"`      //An array of selected Variations for this listing.
}
