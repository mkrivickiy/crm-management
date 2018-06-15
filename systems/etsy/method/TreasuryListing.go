package method

type TreasuryListing struct {
	Data        TreasuryListingData `json:"data"`         //The detailed fields of the listing
	CreationTsz float32             `json:"creation_tsz"` //Time the listing was added to this Treasury, in epoch seconds
}
