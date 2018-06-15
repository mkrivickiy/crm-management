package method

type ListingTranslation struct {
	ListingID   int      `json:"listing_id"`  //The numeric ID for the Listing.
	Language    string   `json:"language"`    //The IETF language tag (e.g. 'fr') for the language of this translation.
	Title       string   `json:"title"`       //The title of the Listing of this Translation.
	Description string   `json:"description"` //The description of the Listing of this Translation.
	Tags        []string `json:"tags"`        //The tags of the Listing of this Translation.
}
