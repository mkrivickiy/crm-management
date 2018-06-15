package method

type TaxonomyNodeProperty struct {
	PropertyID         int                     `json:"property_id"`         //The ID of the property.
	Name               string                  `json:"name"`                //The name of the property.
	DisplayName        string                  `json:"display_name"`        //The name used in user interfaces.
	Scales             []TaxonomyPropertyScale `json:"scales"`              //A list of available scales.
	IsRequired         bool                    `json:"is_required"`         //Is this attribute required for listings in this category?
	SupportsAttributes bool                    `json:"supports_attributes"` //Can this property be used in listing attributes?
	SupportsVariations bool                    `json:"supports_variations"` //Can this property be used in listing variations?
	IsMultivalued      bool                    `json:"is_multivalued"`      //Can this property have multiple values?
	PossibleValues     []TaxonomyPropertyValue `json:"possible_values"`     //A list of allowed values.
	SelectedValues     []TaxonomyPropertyValue `json:"selected_values"`     //A list of values that are always selected for the given category.
}
