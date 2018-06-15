package method

type Treasury struct {
	ID               string            `json:"id"`                 //Unique id of the collection
	Title            string            `json:"title"`              //Title of the collection
	Description      string            `json:"description"`        //Description of the collection
	Homepage         int               `json:"homepage"`           //If the Treasury was featured on the homepage, the time in seconds since the epoch that it was featured
	Mature           bool              `json:"mature"`             //Whether the Treasury has been flagged mature or not
	Private          bool              `json:"private"`            //Whether the Treasury has been set to private
	Locale           string            `json:"locale"`             //Language/Locale of the collection
	CommentCount     int               `json:"comment_count"`      //The number of comments on this Treasury
	Tags             []string          `json:"tags"`               //The tags associated with this Treasury
	Counts           TreasuryCounts    `json:"counts"`             //Clicks, views, shares, and reports metrics of this collection
	Hotness          float32           `json:"hotness"`            //The algorithmic ranking value assigned to this collection
	HotnessColor     string            `json:"hotness_color"`      //The algorithmic ranking value shown as a color (color hexadecimal)
	UserID           int               `json:"user_id"`            //ID of the user (curator) of the collection
	UserName         string            `json:"user_name"`          //Name of the user (curator) of the collection
	UserAvatarID     int               `json:"user_avatar_id"`     //ID of the user's (curator's) avatar
	Listings         []TreasuryListing `json:"listings"`           //The listings that are in this collection
	CreationTsz      float32           `json:"creation_tsz"`       //Time this collection was created, in epoch seconds
	BecamePublicDate int               `json:"became_public_date"` //The time that this treasury was published, in epoch seconds
}
