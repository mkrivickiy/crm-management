package method

type Transaction struct {
	TransactionID    int                           `json:"transaction_id"`     //The numeric ID for this transaction.
	Title            string                        `json:"title"`              //The title of the listing for this transaction.
	Description      string                        `json:"description"`        //The description of the listing for this transaction.
	SellerUserID     int                           `json:"seller_user_id"`     //The numeric ID for the seller of this transaction.
	BuyerUserID      int                           `json:"buyer_user_id"`      //The numeric ID for the buyer of this transaction.
	CreationTsz      float32                       `json:"creation_tsz"`       //The date and time the transaction was created, in epoch seconds.
	PaidTsz          float32                       `json:"paid_tsz"`           //The date and time the transaction was paid, in epoch seconds.
	ShippedTsz       float32                       `json:"shipped_tsz"`        //The date and time the transaction was shipped, in epoch seconds.
	Price            float32                       `json:"price"`              //The price of the transaction.
	CurrencyCode     string                        `json:"currency_code"`      //The ISO code (alphabetic) for the seller's native currency.
	Quantity         int                           `json:"quantity"`           //The quantity of items in this transaction.
	Tags             []string                      `json:"tags"`               //The tags in the listing for this transaction.
	Materials        []string                      `json:"materials"`          //The materials in the listing for this transaction.
	ImageListingID   int                           `json:"image_listing_id"`   //The numeric ID of the primary listing image for this transaction.
	ReceiptID        int                           `json:"receipt_id"`         //The numeric ID for the receipt associated to this transaction.
	ShippingCost     float32                       `json:"shipping_cost"`      //The shipping cost for this transaction.
	IsDigital        bool                          `json:"is_digital"`         //True if this listing is for a digital download.
	FileData         string                        `json:"file_data"`          //Written description of the files attached to this digital listing.
	ListingID        int                           `json:"listing_id"`         //The numeric ID for this listing associated to this transaction.
	IsQuickSale      bool                          `json:"is_quick_sale"`      //True if this transaction was created for an in-person quick sale.
	SellerFeedbackID int                           `json:"seller_feedback_id"` //The numeric ID of seller's feedback.
	BuyerFeedbackID  int                           `json:"buyer_feedback_id"`  //The numeric ID for the buyer's feedback.
	TransactionType  string                        `json:"transaction_type"`   //The type of transaction, usually "listing".
	URL              string                        `json:"url"`                //URL of this transaction
	Variations       []Variations_SelectedProperty `json:"variations"`         //Purchased variations for this transaction.
	ProductData      ListingProduct                `json:"product_data"`       //The product data with the purchased item, if any. Known issue: We have had reports that you may need listings_w permission to access this as the seller. We're investigating this.
}
