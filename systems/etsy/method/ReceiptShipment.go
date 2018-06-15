package method

type ReceiptShipment struct {
	CarrierName       string `json:"carrier_name"`        //Shipping carrier name.
	ReceiptShippingID int    `json:"receipt_shipping_id"` //Receipt shipping id used internally
	TrackingCode      string `json:"tracking_code"`       //Tracking code for carrier.
	TrackingURL       string `json:"tracking_url"`        //Tracking URL for carrier's website.
	BuyerNote         string `json:"buyer_note"`          //Optional note sent to buyer.
	NotificationDate  int    `json:"notification_date"`   //Date the notification was sent.
}
