package method

type FavoriteUser struct {
	UserID         int     `json:"user_id"`          //The user's numeric ID. Note: This field may be absent, depending on the user's privacy settings.
	FavoriteUserID int     `json:"favorite_user_id"` //The numberic ID of this favorite user association. Deprecated: do not use.
	TargetUserID   int     `json:"target_user_id"`   //The targeted favorite user's numeric ID.
	CreationTsz    float32 `json:"creation_tsz"`     //The date and time that the user was favorited.
}
