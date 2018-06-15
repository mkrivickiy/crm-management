package method

type PropertyValue struct {
	PropertyID   int      `json:"property_id"`   //The numeric ID of this property.
	PropertyName string   `json:"property_name"` //The name of the property, in the requested locale language.
	ScaleID      int      `json:"scale_id"`      //The numeric ID of the scale (if any).
	ScaleName    string   `json:"scale_name"`    //The label used to describe the chosen scale (if any).
	ValueIDs     []int    `json:"value_ids"`     //The numeric IDs of the values.
	Values       []string `json:"values"`        //The literal values of the value.
}
