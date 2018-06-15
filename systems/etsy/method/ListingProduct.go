package method

type ListingProduct struct {
	ProductID      int               `json:"product_id"`      //The numeric ID of this product.
	PropertyValues []PropertyValue   `json:"property_values"` //A list of 0-2 properties associated with this product and their values.
	Sku            string            `json:"sku"`             //The product identifier (if set).
	Offerings      []ListingOffering `json:"offerings"`       //A JSON list of active offerings for this product.
	IsDeleted      bool              `json:"is_deleted"`      //Has the product been deleted?
}
