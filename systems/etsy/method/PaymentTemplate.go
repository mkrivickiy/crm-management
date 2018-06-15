package method

type PaymentTemplate struct {
	PaymentTemplateID int    `json:"payment_template_id"` //The numeric ID for this payment template
	AllowBt           bool   `json:"allow_bt"`            //True if the seller accepts bank transfers for payment
	AllowCheck        bool   `json:"allow_check"`         //True if the seller accepts checks for payment
	AllowMo           bool   `json:"allow_mo"`            //True if the seller accepts money order payments
	AllowOther        bool   `json:"allow_other"`         //True if the seller accepts other payments
	AllowPaypal       bool   `json:"allow_paypal"`        //True if the seller accepts paypal
	AllowCc           bool   `json:"allow_cc"`            //True if the seller accepts credit cards
	PaypalEmail       string `json:"paypal_email"`        //The users paypal email address.
	Name              string `json:"name"`                //The name of the seller.
	FirstLine         string `json:"first_line"`          //The first line of the seller's address.
	SecondLine        string `json:"second_line"`         //The second line of the seller's address.
	City              string `json:"city"`                //The seller's city.
	State             string `json:"state"`               //The seller's state.
	Zip               string `json:"zip"`                 //The seller's zip code.
	CountryID         int    `json:"country_id"`          //The seller's country.
	UserID            int    `json:"user_id"`             //The user's numeric ID.
	ListingPaymentID  int    `json:"listing_payment_id"`  //Provided for backward compatibility to ListingPayment. This will return the same value as payment_template_id.
}
