package method

type Feedback struct {
	FeedbackID        int     `json:"feedback_id"`         //The feedback record's numeric ID.
	TransactionID     int     `json:"transaction_id"`      //The transaction's numeric ID.
	CreatorUserID     int     `json:"creator_user_id"`     //The numeric ID of the user who wrote this feedback. Note: This field may be absent, depending on the buyer's privacy settings.
	TargetUserID      int     `json:"target_user_id"`      //The numeric ID of the user who received this feedback. Note: This field may be absent, depending on the buyer's privacy settings.
	SellerUserID      int     `json:"seller_user_id"`      //The numeric ID of the user who was the seller in this transaction.
	BuyerUserID       int     `json:"buyer_user_id"`       //The numeric ID of the user who was the buyer in this transaction. Note: This field may be absent, depending on the buyer's privacy settings.
	CreationTsz       float32 `json:"creation_tsz"`        //Creation time, in epoch seconds.
	Message           string  `json:"message"`             //A message left by the author, explaining the feedback.
	Value             int     `json:"value"`               //The feedback's value; 1 is positive, -1 is negative, and 0 is neutral.
	ImageFeedbackID   int     `json:"image_feedback_id"`   //The numeric ID of the feedback's image (if present). Note: This field may be absent, depending on the buyer's privacy settings.
	ImageURL25x25     string  `json:"image_url_25x25"`     //The url to a photo provided with the feedback, dimensions 25x25. Note: This field may be absent, depending on the buyer's privacy settings.
	ImageURL155x125   string  `json:"image_url_155x125"`   //The url to a photo provided with the feedback, dimensions 155x125. Note: This field may be absent, depending on the buyer's privacy settings.
	ImageURLFullxFull string  `json:"image_url_fullxfull"` //The url to a photo provided with the feedback, dimensions fullxfull. Note: This field may be absent, depending on the buyer's privacy settings.
}
