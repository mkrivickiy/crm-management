package method

type Coupon struct {
	CouponID             int    `json:"coupon_id"`              //The numeric ID of the coupon
	CouponCode           string `json:"coupon_code"`            //The alphanumeric coupon code
	SellerActive         bool   `json:"seller_active"`          //True if the coupon is active
	PctDiscount          int    `json:"pct_discount"`           //The discount percent (null for free shipping coupons)
	FreeShipping         bool   `json:"free_shipping"`          //True if the coupon applies free shipping
	DomesticOnly         bool   `json:"domestic_only"`          //True if the coupon free shipping applies to domestic addresses only
	CurrencyCode         string `json:"currency_code"`          //The 3 letter currency code relating to currency values if any. fixed_discount or minimum_purchase_price
	FixedDiscount        string `json:"fixed_discount"`         //Discount amount the coupon should take. For currency information see currency_code
	MinimumPurchasePrice string `json:"minimum_purchase_price"` //The minimum amount in a cart before tax or shipping needed in order to apply the coupon. For currency information see currency_code
	ExpirationDate       int    `json:"expiration_date"`        //A epoch time (UTC) on which the coupon can no longer be applied
	CouponType           string `json:"coupon_type"`            //The type of coupon: fixed_discount, pct_discount, free_shipping for example
}
