package method

type ShopAboutImage struct {
	ShopID       int    `json:"shop_id"`       //Numeric ID of the shop
	ImageID      int    `json:"image_id"`      //Numeric ID of the image record
	Caption      string `json:"caption"`       //Description of this ShopAboutImage
	Rank         int    `json:"rank"`          //The order of this image in a list
	URL178x178   string `json:"url_178x178"`   //The url to a thumbnail of the image, exactly 178px by 178px.
	URL545xN     string `json:"url_545xN"`     //The url to a thumbnail of the image, no more than 545px width by variable height.
	URL760xN     string `json:"url_760xN"`     //The url to a thumbnail of the image, no more than 760px width by variable height.
	URLFullxFull string `json:"url_fullxfull"` //The url to the full-size image, no more than 1500px width by variable height.
}
