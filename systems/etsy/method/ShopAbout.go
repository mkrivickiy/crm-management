package method

type ShopAbout struct {
	ShopID                int      `json:"shop_id"`                 //Numeric ID of the shop section.
	Status                string   `json:"status"`                  //Status of this data: active or draft. Only Shop owners can see drafts.
	StoryHeadline         string   `json:"story_headline"`          //ShopAbout headline text
	StoryLeadingParagraph string   `json:"story_leading_paragraph"` //The first part of the ShopAbout story
	Story                 string   `json:"story"`                   //The full ShopAbout story text
	RelatedLinks          []string `json:"related_links"`           //The ShopAbout related links
	URL                   string   `json:"url"`                     //URL to the Shop About page on the site.
}
