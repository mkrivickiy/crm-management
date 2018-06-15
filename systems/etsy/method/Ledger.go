package method

type Ledger struct {
	LedgerID   int    `json:"ledger_id"`   //The ledger's numeric ID.
	ShopID     string `json:"shop_id"`     //The shop's numeric ID.
	Currency   string `json:"currency"`    //The currency code of the shop.
	CreateDate int    `json:"create_date"` //The date and time the ledger was created in Epoch seconds.
	UpdateDate int    `json:"update_date"` //The date and time the ledger was last updated in Epoch seconds.
}
