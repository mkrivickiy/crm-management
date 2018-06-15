package method

type Payment struct {
	PaymentID         int    `json:"payment_id"`          //The payment's numeric ID.
	BuyerUserID       int    `json:"buyer_user_id"`       //The buyer's numeric ID.
	ShopID            int    `json:"shop_id"`             //The shop's numeric ID.
	ReceiptID         int    `json:"receipt_id"`          //The receipt's numeric ID.
	AmountGross       int    `json:"amount_gross"`        //The original gross amount of the order, in pennies - this is grand total, including shipping and taxes.
	AmountFees        int    `json:"amount_fees"`         //The original card processing fee of the order in pennies.
	AmountNet         int    `json:"amount_net"`          //The total value of the payment less fees (amount_gross - amount_fees).
	PostedGross       int    `json:"posted_gross"`        //The gross amount which posted to the ledger once shipped. This is equal to the amount_gross UNLESS the seller issues a refund prior to shipping. We consider "shipping" to the event which "posts" to the ledger. Therefore, if the seller refunds first, we simply reduce the amount_gross and post that amount. The seller never sees the refunded amount in their ledger. This is equal to the "Credit" amount in the ledger entry.
	PostedFees        int    `json:"posted_fees"`         //Amount of the fees that posted when shipped. We refund a proportional amount of the fees when a seller refunds a buyer. If they refund prior to shipping, the posted amount will be less then the original.
	PostedNet         int    `json:"posted_net"`          //The total value of the payment at the time of posting it to the ledger less fees (posted_gross - posted_fees)
	AdjustedGross     int    `json:"adjusted_gross"`      //If the payment is refunded, partially or fully, this is the new gross amount after the refund.
	AdjustedFees      int    `json:"adjusted_fees"`       //If the payment is refunded, partially or fully, this is the new fee amount after the refund.
	AdjustedNet       int    `json:"adjusted_net"`        //The total value of the payment after refunds less fees (adjusted_gross - adjusted_fees).
	Currency          string `json:"currency"`            //What currency the payment was made in.
	ShopCurrency      string `json:"shop_currency"`       //The currency of the shop.
	BuyerCurrency     string `json:"buyer_currency"`      //The currency of the buyer.
	ShippingUserID    int    `json:"shipping_user_id"`    //The numeric id of the user to which the order is being shipped.
	ShippingAddressID int    `json:"shipping_address_id"` //The numeric id identifying the shipping address.
	BillingAddressID  int    `json:"billing_address_id"`  //The numeric id identifying the billing address of the buyer.
	Status            string `json:"status"`              //Most commonly "settled" or "authed". Marks the current status of the payment.
	ShippedDate       int    `json:"shipped_date"`        //The date and time the payment was shipped in Epoch seconds.
	CreateDate        int    `json:"create_date"`         //The date and time the payment was created in Epoch seconds.
	UpdateDate        int    `json:"update_date"`         //The date and time the payment was last updated in Epoch seconds.
}
