package method

type Taxonomy struct {
	ID                  int        `json:"id"`                     //The unique ID of this taxonomy node.
	Level               int        `json:"level"`                  //Depth in the taxonomy, roots are at level 0.
	Name                string     `json:"name"`                   //Human readable name for this taxonomy node.
	Parent              string     `json:"parent"`                 //Path of our parent node.
	ParentID            int        `json:"parent_id"`              //ID of our parent node.
	Path                string     `json:"path"`                   //Path from the root to this node, separated by periods.
	CategoryID          int        `json:"category_id"`            //The corresponding ID in Etsy's Category API resource.
	Children            []Taxonomy `json:"children"`               //The child Taxonomy nodes, embedded directly in the response.
	ChildrenIDs         []int      `json:"children_ids"`           //The taxonomy_ids of our child nodes.
	FullPathTaxonomyIDs []int      `json:"full_path_taxonomy_ids"` //The taxonomy_ids of this node and all its parents. They are listed in order from root to leaf.
}
