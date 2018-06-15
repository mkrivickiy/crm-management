package method

type Listing struct {
	ListingID           uint64   `json:"listing_id"`
	State               string   `json:"state,omitempty"`
	UserID              int      `json:"user_id,omitempty"`
	CategoryID          int      `json:"category_id,omitempty"`
	Title               string   `json:"title,omitempty"`
	Description         string   `json:"description,omitempty"`
	CreationTSZ         float32  `json:"creation_tsz,omitempty"`
	EndingTSZ           float32  `json:"ending_tsz,omitempty"`
	OriginalCreationTSZ float32  `json:"original_creation_tsz,omitempty"`
	LastModifiedTSZ     float32  `json:"last_modified_tsz,omitempty"`
	Price               string   `json:"price,omitempty"`
	CurrencyCode        string   `json:"currency_code,omitempty"`
	Quantity            int      `json:"quantity,omitempty"`
	Sku                 []string `json:"sku,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	CategoryPath        []string `json:"category_path,omitempty"`
	CategoryPathIds     []int    `json:"category_path_ids,omitempty"`
	TaxonomyID          int      `json:"taxonomy_id,omitempty"`
	SuggestedTaxonomyID int      `json:"suggested_taxonomy_id,omitempty"`
	TaxonomyPath        []string `json:"taxonomy_path,omitempty"`
	Materials           []string `json:"materials,omitempty"`
	ShopSectionID       int      `json:"shop_section_id,omitempty"`
	FeaturedRank        int      `json:"featured_rank,omitempty"`
	StateTSZ            float32  `json:"state_tsz,omitempty"`
	URL                 string   `json:"url,omitempty"`
	Views               int      `json:"views,omitempty"`
	NumFavorers         int      `json:"num_favorers,omitempty"`
	ShippingTemplateID  int      `json:"shipping_template_id,omitempty"`
	ShippingProfileID   int      `json:"shipping_profile_id,omitempty"`
	ProcessingMin       int      `json:"processing_min,omitempty"`
	ProcessingMax       int      `json:"processing_max,omitempty"`
	WhoMade             string   `json:"who_made,omitempty"`
	IsSupply            string   `json:"is_supply,omitempty"`
	WhenMade            string   `json:"when_made,omitempty"`
	ItemWeight          string   `json:"item_weight,omitempty"`
	ItemWeightUnits     string   `json:"item_weight_units,omitempty"`
	ItemLength          string   `json:"item_length,omitempty"`
	ItemWidth           string   `json:"item_width,omitempty"`
	ItemHeight          string   `json:"item_height,omitempty"`
	ItemDimensionsUnit  string   `json:"item_dimensions_unit,omitempty"`
	IsPrivate           bool     `json:"is_private,omitempty"`
	Recipient           string   `json:"recipient,omitempty"`
	Occasion            string   `json:"occasion,omitempty"`
	Style               []string `json:"style,omitempty"`
	NonTaxable          bool     `json:"non_taxable,omitempty"`
	IsCustomizable      bool     `json:"is_customizable,omitempty"`
	IsDigital           bool     `json:"is_digital,omitempty"`
	FileData            string   `json:"file_data,omitempty"`
	CanWriteInventory   bool     `json:"can_write_inventory,omitempty"`
	HasVariations       bool     `json:"has_variations,omitempty"`
	ShouldAutoRenew     bool     `json:"should_auto_renew,omitempty"`
	Language            string   `json:"language,omitempty"`
}
