package method

type Variations_PropertySet struct {
	PropertySetID        int                                    `json:"property_set_id"`       //Numeric ID of this property set
	Properties           map[int]Variations_PropertySetProperty `json:"properties"`            //Dictionary of properties that are available as variations
	QualifyingProperties map[int]Variations_PropertySetProperty `json:"qualifying_properties"` //Dictionary of properties that are used to add specificity to one or more variation properties
	Options              map[int]string                         `json:"options"`               //Dictionary of available property options for any of the qualifying properties
	Qualifiers           map[int][]Variations_PropertyQualifier `json:"qualifiers"`            //Recursive data-structure indicating the property options available for qualifying_properties, and the nested qualifiers available for each option
}
