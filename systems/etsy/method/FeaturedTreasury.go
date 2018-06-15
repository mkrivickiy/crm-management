package method

type FeaturedTreasury struct {
	TreasuryKey     string  `json:"treasury_key"`      //The string key to identify the Treasury
	TreasuryID      int     `json:"treasury_id"`       //The numeric ID for this FeaturedTreasury.
	TreasuryOwnerID int     `json:"treasury_owner_id"` //The user ID of the FeaturedTreasury owner.
	URL             string  `json:"url"`               //The url to the FeaturedTreasury
	Region          string  `json:"region"`            //The region for which this FeaturedTreasury is active.
	ActiveDate      float32 `json:"active_date"`       //The time this FeaturedTreasury is made active on the homepage
}
