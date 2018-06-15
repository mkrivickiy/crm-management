package method

type LedgerEntry struct {
	LedgerEntryID  int    `json:"ledger_entry_id"` //The ledger entry's numeric ID.
	LedgerID       int    `json:"ledger_id"`       //The ledger's numeric ID.
	Sequence       int    `json:"sequence"`        //The sequence allows ledger entries to be sorted chronologically. The higher the sequence, the more recent the entry.
	CreditAmount   int    `json:"credit_amount"`   //The amount of money credited to the ledger.
	DebitAmount    int    `json:"debit_amount"`    //The amount of money debited from the ledger.
	EntryType      string `json:"entry_type"`      //Details what kind of ledger entry this is: a payment, refund, reversal of a failed refund, disbursement, returned disbursement, recoupment, miscellaneous credit, miscellaneous debit, or bill payment
	ReferenceID    int    `json:"reference_id"`    //Depending on the entry_type, this is the id of the corresponding payment, payment adjustment, or disbursement.
	RunningBalance int    `json:"running_balance"` //The amount of money in the shop's ledger the moment after this entry was applied.
	CreateDate     int    `json:"create_date"`     //The date and time the ledger entry was created in Epoch seconds.
}
