package bill

type Payment struct {
	BillPaymentID int     `json:"bill_payment_id"` //The numeric ID for this bill payment record.
	CreationTsz   float32 `json:"creation_tsz"`    //Creation time, in epoch seconds.
	Type          string  `json:"type"`            //The name of the type of payment.
	TypeID        int     `json:"type_id"`         //The Listing or Transaction ID (or LedgerEntry ID in the case of a ledger payment) to which the payment applies.
	UserID        int     `json:"user_id"`         //The user's numeric ID.
	Amount        float32 `json:"amount"`          //The amount paid.
	CurrencyCode  string  `json:"currency_code"`   //The currency of the payment.
	CreationMonth int     `json:"creation_month"`  //Month that the payment was made.
	CreationYear  int     `json:"creation_year"`   //Year that the payment was made.
}
