package method

type ShopSectionTranslation struct {
	ShopSectionID int    `json:"shop_section_id"` //The numeric ID for the ShopSection.
	Language      string `json:"language"`        //The IETF language tag (e.g. 'fr') for the language of this translation.
	Title         string `json:"title"`           //Translation of title of ShopSection.
}
