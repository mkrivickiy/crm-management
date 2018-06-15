package amocrm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mkrivickiy/crm-management/systems/amocrm/authentificate/response"
)

type RequestInteface interface {
	SetBody(io.Reader) *RequestInteface
	GetBody() io.Reader
	Do() *ResponseInteface
	GetProtocol() string
	GetHost() string
	GetMethodURL() string
	GetRequestURL() string
	GetRequestType() string
	GetContentType() string
}

type Request struct {
	Body        io.Reader
	Host        string
	Protocol    string
	MethodURL   string
	RequestType string
	ContentType string
}

func (r *Request) SetBody(body io.Reader) {
	r.Body = body
}

func (r *Request) GetBody() io.Reader {
	return r.Body
}

func (r *Request) GetProtocol() string {
	return r.Protocol
}

func (r *Request) GetHost() string {
	return r.Host
}

func (r *Request) GetMethodURL() string {
	return r.MethodURL
}

func (r *Request) GetRequestType() string {
	return r.RequestType
}

func (r *Request) GetContentType() string {
	return r.ContentType
}

func (r *Request) GetRequestURL() string {
	return fmt.Sprintf("%s://%s/%s?type=%s", r.GetProtocol(), r.GetHost(), r.GetMethodURL(), r.GetRequestType())
}

func (r *Request) Do() *response.Authentificate {

	request, _ := http.Post(r.GetRequestURL(), r.GetContentType(), r.GetBody())

	defer request.Body.Close()

	response := &response.Authentificate{}

	// rawResponse, _ := ioutil.ReadAll(request.Body)
	// fmt.Println(string(rawResponse))
	_ = json.NewDecoder(request.Body).Decode(response)
	return response
}
