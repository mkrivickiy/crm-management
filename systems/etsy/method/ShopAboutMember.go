package method

type ShopAboutMember struct {
	ShopAboutMemberID    int    `json:"shop_about_member_id"`   //Numeric ID of this ShopAboutMember
	ShopID               int    `json:"shop_id"`                //Numeric ID of the associated Shop
	Name                 string `json:"name"`                   //The name of this ShopAboutMember
	Role                 string `json:"role"`                   //The role of this ShopAboutMember
	Bio                  string `json:"bio"`                    //The bio of this ShopAboutMember
	Rank                 int    `json:"rank"`                   //The order of this member in a list
	IsOwner              bool   `json:"is_owner"`               //Whether this Member is an owner. A Shop can have multiple owners.
	URL90x90             string `json:"url_90x90"`              //The url to a thumbnail of the image, exactly 90px by 90px.
	URL190x190           string `json:"url_190x190"`            //The url to a thumbnail of the image, exactly 190px by 190px.
	HasStaleTranslations bool   `json:"has_stale_translations"` //Whether the shop member translations need to be updated
}
