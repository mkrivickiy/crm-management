package method

type SegmentPoster struct {
	Name      string `json:"name"`       //The formatted name for this SegmentPoster
	Path      string `json:"path"`       //The sequence of slugified names leading to a Segment represented by this SegmentPoster
	ImageURL  string `json:"image_url"`  //A fullxfull image representing this SegmentPoster's content
	ShopID    int    `json:"shop_id"`    //(Optional) The id of the Shop from which the image_url was chosen
	ShopName  int    `json:"shop_name"`  //(Optional) The name of the Shop from which the image_url was chosen
	Weight    int    `json:"weight"`     //A number starting at 1 representing how prominent this SegmentPoster is relative to other SegmentPosters
	ListingID int    `json:"listing_id"` //(Optional) The id of the listing from which the image_url was chosen.
}
