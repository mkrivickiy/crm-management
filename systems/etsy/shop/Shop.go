package shop

type Shop struct {
	ShopID                         int      `json:"shop_id"`                           //The shop's numeric ID.
	ShopName                       string   `json:"shop_name"`                         //The shop's name.
	FirstLine                      string   `json:"first_line"`                        //The first line of the shops address. Deprecated: use UserAddress.first_line instead.
	SecondLine                     string   `json:"second_line"`                       //The second line of the shops address. Deprecated: use UserAddress.second_line instead.
	City                           string   `json:"city"`                              //The city the shop is in. Deprecated: use UserAddress.city instead.
	State                          string   `json:"state"`                             //The state the shop is in. Deprecated: use UserAddress.state instead.
	Zip                            string   `json:"zip"`                               //The zip code the shop is in. Deprecated: use UserAddress.zip instead.
	CountryID                      int      `json:"country_id"`                        //The numeric ID of the country the shop is in. Deprecated: use UserAddress.country_id instead.
	UserID                         int      `json:"user_id"`                           //The numeric ID of the user that owns this shop.
	CreationTsz                    float32  `json:"creation_tsz"`                      //The date and time the shop was created, in epoch seconds.
	Title                          string   `json:"title"`                             //A brief heading for the shop's main page.
	Announcement                   string   `json:"announcement"`                      //An announcement to buyers that displays on the shop's homepage.
	CurrencyCode                   string   `json:"currency_code"`                     //The ISO code (alphabetic) for the seller's native currency.
	IsVacation                     bool     `json:"is_vacation"`                       //If the seller is not currently accepting purchases, 1, blank otherwise.
	CacationMessage                string   `json:"vacation_message"`                  //If is_vacation=1, a message to buyers.
	SaleMessage                    string   `json:"sale_message"`                      //A message that is sent to users who buy from this shop.
	DigitalSaleMessage             string   `json:"digital_sale_message"`              //A message that is sent to users who buy a digital item from this shop.
	LastUpdatedTsz                 float32  `json:"last_updated_tsz"`                  //The date and time the shop was last updated, in epoch seconds.
	ListingActiveCount             int      `json:"listing_active_count"`              //The number of active listings in the shop.
	DigitalListingCount            int      `json:"digital_listing_count"`             //Number of digital listings the shop has.
	LoginName                      string   `json:"login_name"`                        //The user's login name.
	Lat                            float32  `json:"lat"`                               //The latitude of the shop.
	Lon                            float32  `json:"lon"`                               //The longitude of the shop.
	AcceptsCustomRequests          bool     `json:"accepts_custom_requests"`           //True if the shop accepts requests for custom items.
	PolicyWelcome                  string   `json:"policy_welcome"`                    //The introductory text from the top of the seller's policies page (may be blank).
	PolicyPayment                  string   `json:"policy_payment"`                    //The seller's policy on payment (may be blank).
	PolicyShipping                 string   `json:"policy_shipping"`                   //The seller's policy on shipping (may be blank).
	PolicyRefunds                  string   `json:"policy_refunds"`                    //The seller's policy on refunds (may be blank).
	PolicyAdditional               string   `json:"policy_additional"`                 //Any additional policy information the seller provides (may be blank).
	PolicySellerInfo               string   `json:"policy_seller_info"`                //The shop's seller information (may be blank).
	PolicyUpdatedTsz               float32  `json:"policy_updated_tsz"`                //The date and time the shop was last updated, in epoch seconds.
	PolicyHasPrivateReceiptInfo    bool     `json:"policy_has_private_receipt_info"`   //True if seller has private info to show in EU receipts
	CacationAutoreply              string   `json:"vacation_autoreply"`                //If is_vacation=1, a message to buyers in response to new convos.
	GaCode                         string   `json:"ga_code"`                           //The shop's Google Analytics code.
	Name                           string   `json:"name"`                              //The shop owner's real name. Deprecated: use UserAddress.name instead.
	URL                            string   `json:"url"`                               //The URL of the shop
	ImageURLl760x100               string   `json:"image_url_760x100"`                 //The URL to the shop's banner image.
	NumFavorers                    int      `json:"num_favorers"`                      //The number of members who've marked this Shop as a favorite
	Languages                      []string `json:"languages"`                         //The languages that this Shop is enrolled in.
	UpcomingLocalEventID           int      `json:"upcoming_local_event_id"`           //the id of the next upcoming event this hops is attending, that is near the user.
	IconURLFullxFull               string   `json:"icon_url_fullxfull"`                //The url of the shop full-sized shop icon.
	IsUsingStructuredPolicies      bool     `json:"is_using_structured_policies"`      //True if shop has accepted using structured policies.
	HasOnboardedStructuredPolicies bool     `json:"has_onboarded_structured_policies"` //True if shop has viewed structured policies onboarding and accepted OR declined.
	HasUnstructuredPolicies        bool     `json:"has_unstructured_policies"`         //True if shop has any unstructured policy fields filled out.
	PolicyPrivacy                  string   `json:"policy_privacy"`                    //Privacy policy information the seller provides (may be blank).
	UseNewInventoryEndpoints       bool     `json:"use_new_inventory_endpoints"`       //Should this user's listings be created or edited using the new Inventory endpoints?
	IncludeDisputeFormLink         bool     `json:"include_dispute_form_link"`         //True if shop policies include a link to EU online dispute form
}
