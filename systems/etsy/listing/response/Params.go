package response

type Params struct {
	Limit             int      `json:"limit"`
	Offset            int      `json:"offset"`
	Page              int      `json:"page"`
	Keywords          string   `json:"keywords"`
	SortOn            string   `json:"sort_on"`
	SortOrder         string   `json:"sort_order"`
	MinPrice          float32  `json:"min_price"`
	MaxPrice          float32  `json:"max_price"`
	Color             string   `json:"color"`
	ColorAccuracy     int      `json:"color_accuracy"`
	Tags              []string `json:"tags"`
	Category          string   `json:"category"`
	Location          string   `json:"location"`
	Lat               float32  `json:"lat"`
	Lon               float32  `json:"lon"`
	Region            string   `json:"region"`
	GeoLevel          string   `json:"geo_level"`
	AcceptsGiftCards  bool     `json:"accepts_gift_cards"`
	TranslateKeywords bool     `json:"translate_keywords"`
}
