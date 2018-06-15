package method

type FavoriteListing struct {
	ListingID    int    `json:"listing_id"`    //The listings numeric ID.
	UserID       int    `json:"user_id"`       //The user's numeric ID. Note: This field may be absent, depending on the user's privacy settings.
	ListingState string `json:"listing_state"` //The state of the listing.
	CreateDate   int    `json:"create_date"`   //The date and time that the listing was favorited.
}
