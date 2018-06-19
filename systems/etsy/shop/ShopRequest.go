package shop

import (
	"encoding/json"
	"net/http"

	"github.com/mkrivickiy/crm-management/systems/etsy/base"
)

type ShopRequest struct {
	base.Request
}

func NewRequest() *ShopRequest {
	request := ShopRequest{}
	request.Host = "openapi.etsy.com"
	request.Protocol = "https"
	request.MethodURL = "v2/users/20191168/shops"
	request.ApiKey = "qytwegngo28vg8hh1k8sfzra"
	request.ContentType = "application/json; charset=utf-8"
	return &request
}

func (r *ShopRequest) Get() *ShopResponse {

	request, _ := http.Get(r.GetRequestURL())

	defer request.Body.Close()

	response := &ShopResponse{}

	_ = json.NewDecoder(request.Body).Decode(response)
	return response
}
