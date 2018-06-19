package content

import (
	"net/url"
)

type Values struct {
}

func (v *Values) Encode() string {
	val := url.Values{}
	val.Set("name", "Ava")
	val.Add("friend", "Jess")
	val.Add("friend", "Sarah")
	val.Add("friend", "Zoe")
	encodedContent := val.Encode()
	return encodedContent
}
