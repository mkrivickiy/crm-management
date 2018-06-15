package method

type User struct {
	UserID                   int          `json:"user_id"`                     //The user's numeric ID. This is also valid as the user's shop ID.
	LoginName                string       `json:"login_name"`                  //The user's login name.
	PrimaryEmail             string       `json:"primary_email"`               //The user's primary email address.
	CreationTsz              float32      `json:"creation_tsz"`                //The date and time the user was created, in epoch seconds.
	UserPubKey               string       `json:"user_pub_key"`                //Public key for user
	ReferredByUserID         int          `json:"referred_by_user_id"`         //The numeric ID of the user that referred this user.
	FeedbackInfo             FeedbackInfo `json:"feedback_info"`               //An associative array of feedback totals for the user.
	AwaitingFeedbackCount    int          `json:"awaiting_feedback_count"`     //The total number of transactions the user has available for a new review.
	UseNewInventoryEndpoints bool         `json:"use_new_inventory_endpoints"` //Should this user's listings be created or edited using the new Inventory endpoints?
}
