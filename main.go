package main

import (
	"fmt"

	"github.com/mkrivickiy/crm-management/core/rest/content"
)

func main() {

	// request := authentificate.NewRequest()
	// data := authentificate.Authentificate{Username: "mkrivickiy@gmail.com", Userhash: "470710cae5dedc3a48713b0f07145125"}
	// request.SetBody(data.Encode())
	// result := request.Do()
	// fmt.Println(result)

	// request := listing.NewRequest()
	// result := request.Get()
	// fmt.Println(*result)

	// request := request.Post{}
	content := content.Values{}
	encodedString := content.Encode()
	fmt.Println("Hello " + encodedString)
}
