package method

type Variations_PropertyQualifier struct {
	PropertyID int                                     `json:"property_id"` //The numeric ID of this property
	Options    []int                                   `json:"options"`     //An array of numeric IDs of available options for this property qualifier
	Results    map[string]Variations_PropertyQualifier `json:"results"`     //Variations_PropertyQualifier)	Recursive qualifiers representative of inreased property specificity (Keys are numeric strings)
	Aliases    map[string]int                          `json:"aliases"`     //Option values listed here share qualifier settings with other options (Keys are numeric strings)
}
