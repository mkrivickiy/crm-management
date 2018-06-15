package method

type ShippingTemplate struct {
	ShippingTemplateID         int    `json:"shipping_template_id"`          //The numeric ID of this shipping template.
	Title                      string `json:"title"`                         //The name of this shipping template.
	UserID                     int    `json:"user_id"`                       //The numeric ID of the user who owns this shipping template.
	MinProcessingDays          int    `json:"min_processing_days"`           //The minimum number of days for processing the listing.
	MaxProcessingDays          int    `json:"max_processing_days"`           //The maximum number of days for processing the listing.
	ProcessingDaysDisplayLabel string `json:"processing_days_display_label"` //Translated display label for processing days.
	OriginCountryID            int    `json:"origin_country_id"`             //The numeric ID of the country from which the listing ships.
}
