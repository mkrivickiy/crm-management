package amocrm

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type Auth struct {
	Username string `json:"USER_LOGIN"`
	Userhash string `json:"USER_HASH"`
}

func (a *Auth) Encode() string {
	values := url.Values{}
	values.Add("USER_LOGIN", a.Username)
	values.Add("USER_HASH", a.Userhash)

	return values.Encode()
}

func (a *Auth) EncodeJson() *bytes.Buffer {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(a)
	return buffer
}
