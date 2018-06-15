package method

type Variations_PropertySetProperty struct {
	PropertyID    int    `json:"property_id"`    //This numeric ID of this property
	Name          string `json:"name"`           //The name of this property
	InputName     string `json:"input_name"`     //Suggested form field name
	Label         string `json:"label"`          //Descriptive text label for form input
	Param         string `json:"param"`          //Specifies the qualifier parameter when requesting suggested options for a property
	DefaultOption string `json:"default_option"` //Initial option value for form input
	Description   string `json:"description"`    //Text description of property
	IsCustom      bool   `json:"is_custom"`      //True if this property is Custom 1 or Custom 2
	Title         string `json:"title"`          //Property name, formatted for display
}
