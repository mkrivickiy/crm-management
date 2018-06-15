package method

type Money struct {
	Amount               int    `json:"amount"`                 //The amount of money represented by this data.
	Divisor              int    `json:"divisor"`                //The divisor to render the amount
	CurrencyCode         string `json:"currency_code"`          //The requested locale currency.
	FormattedRaw         string `json:"formatted_raw"`          //The formatted amount without codes or symbols in the requested locale's numeric style, e.g. '10.42'.
	FormattedShort       string `json:"formatted_short"`        //The formatted amount with a symbol in the requested locale's numeric style, e.g. 'US$10.42'.
	FormattedLong        string `json:"formatted_long"`         //The formatted amount with a symbol and currency in the requested locale's numeric style, e.g. '$10.42 USD'.
	OriginalCurrencyCode string `json:"original_currency_code"` //The original currency code the value was listed in (if the value has been converted). Deprecated: Replaced by "before_conversion" (to be removed 15 February 2017).
	BeforeConversion     *Money `json:"before_conversion"`      //A representation of the value without currency conversion (if conversion has happened).
}
