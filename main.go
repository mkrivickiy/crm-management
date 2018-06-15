package main

import (
	"fmt"

	"github.com/mkrivickiy/crm-management/systems/amocrm/authentificate"
)

func main() {

	request := authentificate.NewRequest()
	data := authentificate.Authentificate{Username: "mkrivickiy@gmail.com", Userhash: "470710cae5dedc3a48713b0f07145125"}
	request.SetBody(data.Encode())
	result := request.Do()
	fmt.Println(result)
}
