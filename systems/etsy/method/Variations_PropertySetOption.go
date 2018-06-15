package method

type Variations_PropertySetOption struct {
	PropertyOptionID int    `json:"property_option_id"` //The property option's unique ID
	Name             string `json:"name"`               //property option name
	FormattedName    string `json:"formatted_name"`     //property option name, formatted for display
}
