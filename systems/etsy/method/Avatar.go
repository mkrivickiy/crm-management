package method

type Avatar struct {
	AvatarID        int     `json:"avatar_id"`          //The numeric ID for this avatar image.
	HexCode         string  `json:"hex_code"`           //The avatar' average RGB color, in webhex notation.
	Red             int     `json:"red"`                //The avatar's average red channel (RGB color) value from 0-255.
	Green           int     `json:"green"`              //The avatar's average green channel (RGB color) value from 0-255.
	Blue            int     `json:"blue"`               //The avatar's average blue channel (RGB color) value from 0-255.
	Hue             int     `json:"hue"`                //The avatar's average hue (HSV color) from 0-360.
	Saturation      int     `json:"saturation"`         //The avatar's average saturation (HSV color) from 0-100.
	Brightness      int     `json:"brightness"`         //The avatar's average brightness (HSV color) from 0-100.
	IsBlackAndWhite bool    `json:"is_black_and_white"` //True if the avatar is a black and white image.
	CreationTsz     float32 `json:"creation_tsz"`       //The time that the avatar was uploaded.
	UserID          int     `json:"user_id"`            //The numeric ID of the user who owns the avatar.
}
