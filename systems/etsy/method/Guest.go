package method

type Guest struct {
	GuestID     string `json:"guest_id"`     //The ID that associates to the guest - a random 13-digit hash. Validated by the regular expression '/^[a-zA-Z0-9-_]+$/'. Each ID is identified internally to the Application that created it, meaning that other applications will not have access to this guest, even with the same ID.
	CheckoutURL string `json:"checkout_url"` //The URL to send the user's browser to on checkout. This will allow them to use this guest ID to check out with the carts associated. Please note that a guest cart's data are undefined after being merged on checkout.
}
