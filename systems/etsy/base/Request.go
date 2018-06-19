package base

import (
	"fmt"
	"io"
)

type RequestInteface interface {
	SetBody(io.Reader) *RequestInteface
	GetBody() io.Reader
	Do() *ResponseInteface
	GetProtocol() string
	GetHost() string
	GetMethodURL() string
	GetRequestURL() string
	GetApiKey() string
	GetContentType() string
}

type Request struct {
	Body        io.Reader
	Host        string
	Protocol    string
	MethodURL   string
	ApiKey      string
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

func (r *Request) GetApiKey() string {
	return r.ApiKey
}

func (r *Request) GetContentType() string {
	return r.ContentType
}

func (r *Request) GetRequestURL() string {
	return fmt.Sprintf("%s://%s/%s?api_key=%s", r.GetProtocol(), r.GetHost(), r.GetMethodURL(), r.GetApiKey())
}
