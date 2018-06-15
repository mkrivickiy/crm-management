package method

type Variations_Property struct {
	PropertyID    int                 `json:"property_id"`    //This numeric ID of this property
	FormattedName string              `json:"formatted_name"` //This translated name of this property
	Options       []Variations_Option `json:"options"`        //Available options for this property
}
