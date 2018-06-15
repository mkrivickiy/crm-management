package method

type TaxonomyPropertyValue struct {
	ValueID int    `json:"value_id"` //The numeric ID of this value.
	Name    string `json:"name"`     //How to display the value.
	ScaleID int    `json:"scale_id"` //The ID of the scale this belongs to (if any).
	Order   int    `json:"order"`    //The order of this value under its property
	EqualTo []int  `json:"equal_to"` //The list of values this value is equal to (if any).
}
