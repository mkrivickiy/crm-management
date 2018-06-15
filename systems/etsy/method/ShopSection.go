package method

type ShopSection struct {
	ShopSectionID      int    `json:"shop_section_id"`      //The numeric ID of the shop section.
	Title              string `json:"title"`                //The title of the section.
	Rank               int    `json:"rank"`                 //Display order.
	UserID             int    `json:"user_id"`              //The ID of the user who owns this shop section.
	ActiveListingCount int    `json:"active_listing_count"` //The number of active listings currently in the section.
}
