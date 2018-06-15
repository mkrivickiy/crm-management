package method

type Receipt struct {
	ReceiptID                int               `json:"receipt_id"`                 //The receipt's numeric ID.
	ReceiptType              int               `json:"receipt_type"`               //The enum of the order type this receipt is associated with.
	OrderID                  int               `json:"order_id"`                   //The numeric ID of the order this receipt is associated with.
	SellerUserID             int               `json:"seller_user_id"`             //The numeric ID of the seller for this receipt.
	BuyerUserID              int               `json:"buyer_user_id"`              //The numeric ID of the buyer for this receipt.
	CreationTsz              float32           `json:"creation_tsz"`               //Creation time, in epoch seconds.
	LastModifiedTsz          float32           `json:"last_modified_tsz"`          //Last modification time, in epoch seconds.
	Name                     string            `json:"name"`                       //The name portion of the buyer's address.
	FirstLine                string            `json:"first_line"`                 //The first line of the buyer's address.
	SecondLine               string            `json:"second_line"`                //The second line of the buyer's address.
	City                     string            `json:"city"`                       //The city for the buyer's address.
	State                    string            `json:"state"`                      //The state for the buyer's address.
	Zip                      string            `json:"zip"`                        //The zip code of the buyer's address.
	FormattedAddress         string            `json:"formatted_address"`          //The locally formatted address strng of the buyer's shipping address.
	CountryID                int               `json:"country_id"`                 //The numeric ID of the buyer's country.
	PaymentMethod            string            `json:"payment_method"`             //The payment method used. May be pp, cc, ck, mo, or other (stands for paypal, credit card, check, money order, other).
	PaymentEmail             string            `json:"payment_email"`              //The email address where payment confirmation is sent.
	MessageFromSeller        string            `json:"message_from_seller"`        //An optional message from the seller.
	MessageFromBuyer         string            `json:"message_from_buyer"`         //An optional message from the buyer.
	WasPaid                  bool              `json:"was_paid"`                   //True if the receipt was paid.
	TotalTaxCost             float32           `json:"total_tax_cost"`             //The total sales tax of the receipt.
	TotalVatCost             float32           `json:"total_vat_cost"`             //The total VAT of the receipt.
	TotalPrice               float32           `json:"total_price"`                //Sum of the individual listings' (price * quantity). Does not included tax or shipping costs.
	TotalShippingCost        float32           `json:"total_shipping_cost"`        //The total shipping cost of the receipt.
	CurrencyCode             string            `json:"currency_code"`              //The ISO code (alphabetic) for the seller's native currency.
	MessageFromPayment       string            `json:"message_from_payment"`       //The machine generated acknowledgement from the payment system.
	WasShipped               bool              `json:"was_shipped"`                //True if the items in the receipt were shipped.
	BuyerEmail               string            `json:"buyer_email"`                //The email address of the buyer.
	SellerEmail              string            `json:"seller_email"`               //The email address of the seller.
	IsGift                   bool              `json:"is_gift"`                    //Whether the buyer marked item as a gift.
	NeedsGiftWrap            bool              `json:"needs_gift_wrap"`            //Whether the buyer purchased gift wrap.
	GiftMessage              string            `json:"gift_message"`               //The message the buyer wants sent with the gift.
	GiftWrapPrice            float32           `json:"gift_wrap_price"`            //The gift wrap price of the receipt.
	DiscountAmount           float32           `json:"discount_amt"`               //The total discount for the receipt, if using a discount (percent or fixed) Coupon. Free shipping Coupons are not reflected here; check the Coupon object for these.
	Subtotal                 float32           `json:"subtotal"`                   //total_price minus coupon discounts. Does not included tax or shipping costs.
	Grandtotal               float32           `json:"grandtotal"`                 //total_price minus coupon discount plus tax and shipping costs.
	AdjustedGrandtotal       float32           `json:"adjusted_grandtotal"`        //grandtotal after payment adjustments.
	ShippingTrackingCode     string            `json:"shipping_tracking_code"`     //The tracking code for the shipment. Deprecated: This field will be removed by 10/2013. Use shipments instead. When a receipt has more than one shipment, behavior of this field is unspecified.
	ShippingTrackingURL      string            `json:"shipping_tracking_url"`      //The tracking URL for the shipment. Deprecated: This field will be removed by 10/2013. Use shipments instead. When a receipt has more than one shipment, behavior of this field is unspecified.
	ShippingCarrier          string            `json:"shipping_carrier"`           //The shipping carrier. Deprecated: This field will be removed by 10/2013. Use shipments instead. When a receipt has more than one shipment, behavior of this field is unspecified.
	ShippingNote             string            `json:"shipping_note"`              //The shipping notification note text. Deprecated: This field will be removed by 10/2013. Use shipments instead. When a receipt has more than one shipment, behavior of this field is unspecified.
	ShippingNotificationDate int               `json:"shipping_notification_date"` //The date the last shipping notification was sent. Deprecated: This field will be removed by 10/2013. Use shipments instead. When a receipt has more than one shipment, behavior of this field is unspecified.
	Shipments                []ReceiptShipment `json:"shipments"`                  //Shipment information associated to this receipt.
}
