package shop

import (
	"github.com/mkrivickiy/crm-management/systems/etsy/shop/response"
)

type ShopResponse struct {
	Count      int                 `json:"count"`
	Pagination response.Pagination `json:"pagination"`
	Params     response.Params     `json:"params"`
	Result     []Shop              `json:"results"`
	Type       string              `json:"type"`
}
