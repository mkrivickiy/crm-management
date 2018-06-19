package listing

import (
	"encoding/json"
	"net/http"

	"github.com/mkrivickiy/crm-management/systems/etsy/base"
)

type ListingRequest struct {
	base.Request
}

func NewRequest() *ListingRequest {
	request := ListingRequest{}
	request.Host = "openapi.etsy.com"
	request.Protocol = "https"
	request.MethodURL = "v2/listings/active"
	request.ApiKey = "qytwegngo28vg8hh1k8sfzra"
	request.ContentType = "application/json; charset=utf-8"
	return &request
}

func (r *ListingRequest) Get() *ListingResponse {

	request, _ := http.Get(r.GetRequestURL())

	defer request.Body.Close()

	response := &ListingResponse{}

	_ = json.NewDecoder(request.Body).Decode(response)
	return response
}
