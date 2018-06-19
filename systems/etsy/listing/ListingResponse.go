package listing

import (
	"github.com/mkrivickiy/crm-management/systems/etsy/listing/response"
)

type ListingResponse struct {
	Count      int                 `json:"count"`
	Pagination response.Pagination `json:"pagination"`
	Params     response.Params     `json:"params"`
	Result     []Listing           `json:"results"`
	Type       string              `json:"type"`
}
