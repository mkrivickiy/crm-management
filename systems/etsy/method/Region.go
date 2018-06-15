package method

type Region struct {
	RegionID   int    `json:"region_id"`   //The numeric ID of this region.
	RegionName string `json:"region_name"` //The name of the region.
	IsDead     bool   `json:"is_dead"`     //The eligibilty of this region to be used
}
