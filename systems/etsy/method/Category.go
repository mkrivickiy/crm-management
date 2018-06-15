package method

type Category struct {
	CategoryID      int    `json:"category_id"`      //The identifier for this category.
	Name            string `json:"name"`             //The programmatic name for this category.
	MetaTitle       string `json:"meta_title"`       //The "title" meta tag value for the category's landing page (may be null).
	MetaKeywords    string `json:"meta_keywords"`    //The "keywords" meta tag value for the category's landing page (may be null).
	MetaDescription string `json:"meta_description"` //The "description" meta tag value for the category's landing page (may be null).
	PageDescription string `json:"page_description"` //A short description of the category from the category' landing page (may be null).
	PageTitle       string `json:"page_title"`       //The title of the category's landing page (may be null).
	CategoryName    string `json:"category_name"`    //The category's string ID.
	ShortName       string `json:"short_name"`       //A short display name for the category.
	LongName        string `json:"long_name"`        //A longer display name for the category.
	NumChildren     int    `json:"num_children"`     //The number of subcategories below this one. Subcatgories append a new tag to the end of their parent's category_name.
}
