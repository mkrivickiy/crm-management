package method

type ListingImage struct {
	ListingImageID  int     `json:"listing_image_id"`   //The numeric ID of the listing image.
	HexCode         string  `json:"hex_code"`           //The image's average color, in webhex notation.
	Red             int     `json:"red"`                //The image's average red value, 0-255 (RGB color).
	Green           int     `json:"green"`              //The image's average green value, 0-255 (RGB color).
	Blue            int     `json:"blue"`               //The image's average blue value, 0-255 (RGB color).
	Hue             int     `json:"hue"`                //The image's average hue, 0-360 (HSV color).
	Saturation      int     `json:"saturation"`         //The image's average saturation, 0-100 (HSV color).
	Brightness      int     `json:"brightness"`         //The image's average brightness, 0-100 (HSV color).
	IsBlackAndWhite bool    `json:"is_black_and_white"` //True if the image is in black & white.
	CreationTsz     float32 `json:"creation_tsz"`       //Creation time, in epoch seconds.
	ListingID       int     `json:"listing_id"`         //The numeric value of the listing id the image belongs to.
	Rank            int     `json:"rank"`               //Display order.
	URL75x75        string  `json:"url_75x75"`          //The url to a 75x75 thumbnail of the image.
	URL170x135      string  `json:"url_170x135"`        //The url to a 170x135 thumbnail of the image.
	URL570xN        string  `json:"url_570xN"`          //The url to a thumbnail of the image, no more than 570px width by variable height.
	URLFullxFull    string  `json:"url_fullxfull"`      //The url to the full-size image, up to 3000px in each dimension.
	Fullheight      int     `json:"full_height"`        //Height of the image returned by url_fullxfull.
	Fullwidth       int     `json:"full_width"`         //Width of the image returned by url_fullxfull.
}
