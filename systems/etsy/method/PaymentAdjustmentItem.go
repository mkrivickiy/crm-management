package method

type PaymentAdjustmentItem struct {
	PaymentAdjustmentItemID int    `json:"payment_adjustment_item_id"` //The payment adjustment item's numeric ID.
	PaymentAdjustmentID     int    `json:"payment_adjustment_id"`      //The payment adjustment's numeric ID.
	AdjustmentType          string `json:"adjustment_type"`            //A human readable string describing the area being refundedOne of: TRANS, TAX, or SHIP
	Amount                  int    `json:"amount"`                     //The total amount of the payment adjustment item.
	TransactionID           int    `json:"transaction_id"`             //The numerice ID of the Credit Card Transaction
	CreateDate              int    `json:"create_date"`                //The date and time the payment adjustment item was created in Epoch seconds.
}
