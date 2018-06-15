package method

type Manufacturer struct {
	Name        string   `json:"name"`        //Name of the manufacturer
	Description []string `json:"description"` //The manufacturer's description as it appears on the Shop About page.
	Location    string   `json:"location"`    //The manufacturer's location.
}
