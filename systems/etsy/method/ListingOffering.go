package method

type ListingOffering struct {
	OfferingID int   `json:"offering_id"` //The numeric ID of this offering.
	Price      Money `json:"price"`       //The price of the product
	Quantity   int   `json:"quantity"`    //How many of this product are available?
	IsEnabled  bool  `json:"is_enabled"`  //Is the offering shown to buyers?
	IsDeleted  bool  `json:"is_deleted"`  //Has the offering been deleted?
}
