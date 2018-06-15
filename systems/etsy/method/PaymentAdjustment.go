package method

type PaymentAdjustment struct {
	PaymentAdjustmentID        int    `json:"payment_adjustment_id"`         //The payment adjustment's numeric ID.
	PaymentID                  int    `json:"payment_id"`                    //The payment's numeric ID.
	Status                     string `json:"status"`                        //The status of the Payment Adjustment. Can be OPEN, REFUNDED, ERROR, or REFUND_FAILED.
	IsSuccess                  bool   `json:"is_success"`                    //Whether the Payment Adjustment was successful or will likely be completed successfully.
	UserID                     int    `json:"user_id"`                       //The seller's numeric ID.
	ReasonCode                 string `json:"reason_code"`                   //A human-readable string describing the need for the refund.
	TotalAdjustmentAmount      int    `json:"total_adjustment_amount"`       //The total amount of the refund in the payment currency.
	ShopTotalAdjustmentAmount  int    `json:"shop_total_adjustment_amount"`  //The total amount of the refund in the shop currency.
	BuyerTotalAdjustmentAmount int    `json:"buyer_total_adjustment_amount"` //The total amount of the refund in the buyer currency.
	TotalFeeAdjustmentAmount   int    `json:"total_fee_adjustment_amount"`   //The amount of card processing fees associated with this adjustment.
	CreateDate                 int    `json:"create_date"`                   //The date and time the payment adjustment was created in Epoch seconds.
	UpdateDate                 int    `json:"update_date"`                   //The date and time the payment adjustment was last updated in Epoch seconds.
}
