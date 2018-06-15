package method

type ShopTranslation struct {
	ShopID             int    `json:"shop_id"`              //The numeric ID for the Shop.
	Language           string `json:"language"`             //`json:"shop_section_id"` //The IETF language tag (e.g. 'fr') for the language of this translation.
	Announcement       string `json:"announcement"`         //Translation of an announcement to buyers that displays on the shop's homepage.
	PolicyWelcome      string `json:"policy_welcome"`       //Translation of the introductory text from the top of the seller's policies page (may be blank).
	PolicyPayment      string `json:"policy_payment"`       //Translation of the seller's policy on payment (may be blank).
	PolicyShipping     string `json:"policy_shipping"`      //Translation of the seller's policy on shipping (may be blank).
	PolicyRefunds      string `json:"policy_refunds"`       //Translation of the seller's policy on refunds (may be blank).
	PolicyAdditional   string `json:"policy_additional"`    //Translation of any additional policy information the seller provides (may be blank).
	PolicyPrivacy      string `json:"policy_privacy"`       //Translation of privacy policy information the seller provides (may be blank).
	PolicySellerInfo   string `json:"policy_seller_info"`   //Translation of the Shop's Seller information (may be blank).
	SaleMessage        string `json:"sale_message"`         //Translation of a message that is sent to users who buy from this shop.
	DigitalSaleMessage string `json:"digital_sale_message"` //Translation of a message that is sent to users who buy a digital item from this shop.
	Title              string `json:"title"`                //Translation of a brief heading for the shop's main page.
	VacationAutoreply  string `json:"vacation_autoreply"`   //Translation of a message to buyers in response to new convos (if Shop.is_vacation is true).
	VacationMessage    string `json:"vacation_message"`     //Translation of a message to buyers (if Shop.is_vacation is true).
}
