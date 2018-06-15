package authentificate

import (
	"bytes"
	"encoding/json"
)

type Authentificate struct {
	Username string `json:"USER_LOGIN"`
	Userhash string `json:"USER_HASH"`
}

func (a *Authentificate) Encode() *bytes.Buffer {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(a)
	return buffer
}
