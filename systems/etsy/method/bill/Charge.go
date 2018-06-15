package bill

type Charge struct {
	BillChargeID    int     `json:"bill_charge_id"`    //The numeric ID for this bill charge record.
	CreationTsz     float32 `json:"creation_tsz"`      //Creation time, in epoch seconds.
	Type            string  `json:"type"`              //The name of the type of charge.
	TypeID          int     `json:"type_id"`           //The Listing, Transaction or Shipping Label ID to which the charge applies.
	UserID          int     `json:"user_id"`           //The user's numeric ID.
	Amount          float32 `json:"amount"`            //The amount charged.
	CurrencyCode    string  `json:"currency_code"`     //The currency of the charge.
	CreationYear    int     `json:"creation_year"`     //Year that the charge was created.
	CreationMonth   int     `json:"creation_month"`    //Month that the charge was created.
	LastModifiedTsz float32 `json:"last_modified_tsz"` //Time when charge was last modified.
}
