package method

type DataType struct {
	Type   string   `json:"type"`   //Base type of data
	Values []string `json:"values"` //Allowable values (for an enum.)
}
