package method

type ListingInventory struct {
	Products           []ListingProduct `json:"products"`             //The products available for this listing.
	PriceOnProperty    []int            `json:"price_on_property"`    //Which properties control price?
	QuantityOnProperty []int            `json:"quantity_on_property"` //Which properties control quantity?
	SkuOnProperty      []int            `json:"sku_on_property"`      //Which properties control SKU?
}
