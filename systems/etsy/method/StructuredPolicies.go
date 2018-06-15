package method

type StructuredPolicies struct {
	StructuredPoliciesID         int    `json:"structured_policies_id"` //Unqiue identifier for the policies
	Payments                     string `json:"payments"`               //Payment information for the seller.
	Refunds                      string `json:"refunds"`                //Structured refunds information for the seller.
	Shipping                     string `json:"shipping"`               //Structured shipping information for the seller.
	CreateDate                   int    `json:"create_date"`
	UpdateDate                   int    `json:"update_date"`
	CreateDateFormatted          string `json:"create_date_formatted"`
	UpdateDateFormatted          string `json:"update_date_formatted"`
	HasUnstructuredPolicies      bool   `json:"has_unstructured_policies"` //True if the shop has previously filled out any unstructured policies.
	AdditionalTermsAndConditions string `json:"additional_terms_and_conditions"`
	ShopInGermany                bool   `json:"shop_in_germany"`           //True if the shop is in Germany
	CharLimits                   string `json:"char_limits"`               //Character limits for custom fields
	IncludeDisputeFormLink       bool   `json:"include_dispute_form_link"` //Whether to include a link to dispute resolution form.
	Privacy                      string `json:"privacy"`                   //Privacy policy for the shop.
}
