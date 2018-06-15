package method

type Variations_SelectedProperty struct {
	PropertyID     int    `json:"property_id"`     //The numeric ID of the selected property
	ValueID        int    `json:"value_id"`        //The numeric ID of selected value
	FormattedName  string `json:"formatted_name"`  //The formatted/translated name of the selected property
	FormattedValue string `json:"formatted_value"` //The formatted/translated name of the selected value
	IsValid        bool   `json:"is_valid"`        //Whether the selected variation value is a valid option for this property. NOTE: This field will not be present on Transactions.
}
