package request

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mkrivickiy/crm-management/core/rest"
)

type Post struct {
	rest.Request
}

func (p *Post) Do() rest.ResponseInterface {

	config := rest.Config{URL: "", ContentType: ""} //p.GetConfig()
	// val := content.Values{}
	// p.SetContent(val)
	content := p.GetContent()
	response, err := http.Post(config.GetURL(), config.GetContentType(), strings.NewReader(content.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	return p.GetResponse()
}
