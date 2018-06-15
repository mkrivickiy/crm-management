package method

type TaxonomyPropertyScale struct {
	ScaleID     int    `json:"scale_id"`     //The ID of the scale.
	DisplayName string `json:"display_name"` //How to label the scale.
	Description string `json:"description"`  //A description of the scale.
}
