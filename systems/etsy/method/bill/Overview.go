package bill

import "time"

type Overview struct {
	IsOverdue      bool      `json:"is_overdue"`      //True if the user has an overdue balance.
	CurrencyCode   string    `json:"currency_code"`   //The currency unit in which the user is billed.
	OverdueBalance float32   `json:"overdue_balance"` //The total overdue balance owed by the user.
	BalanceDue     float32   `json:"balance_due"`     //The total amount currently due for payment (including any overdue balance.)
	TotalBalance   float32   `json:"total_balance"`   //The total amount owed by the user (including fees that are not yet due for payment.)
	DateDue        time.Time `json:"date_due"`        //The date by which the user's balance due must be paid.
	DateOverdue    time.Time `json:"date_overdue"`    //The date on which the user's balance became overdue.
}
