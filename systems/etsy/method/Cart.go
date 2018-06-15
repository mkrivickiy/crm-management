package method

type Cart struct {
	CartID                 int            `json:"cart_id"`                  //The numeric ID of the cart
	ShopName               string         `json:"shop_name"`                //The shop name
	Message                string         `json:"message_to_seller"`        //The message to the seller
	DestinationCountryID   int            `json:"destination_country_id"`   //The numeric ID of the destination country
	CouponCode             string         `json:"coupon_code"`              //The alphanumeric coupon code applied to the cart. Deprecated: The Etsy API no longer supports coupons.
	CurrencyCode           string         `json:"currency_code"`            //The ISO (alphabetic) code for the currency
	Total                  string         `json:"total"`                    //The total price
	Subtotal               string         `json:"subtotal"`                 //The subtotal price
	ShippingCost           string         `json:"shipping_cost"`            //The shipping cost
	TaxCost                string         `json:"tax_cost"`                 //The tax cost
	DiscountAmount         string         `json:"discount_amount"`          //The line-item discount amount (does not include tax or shipping)
	ShippingDiscountAmount string         `json:"shipping_discount_amount"` //The shipping discount amount
	TaxDiscountAmount      string         `json:"tax_discount_amount"`      //The tax discount amount
	URL                    string         `json:"url"`                      //The full URL to the cart page on Etsy
	Listings               []CartListing  `json:"listings"`                 //An array of purchase information for the listings
	IsDownloadOnly         bool           `json:"is_download_only"`         //The cart is download only
	HasVat                 bool           `json:"has_vat"`                  //The cart has VAT tax
	ShippingOption         ShippingOption `json:"shipping_option"`          //The selected shipping option identifier for the cart
}
