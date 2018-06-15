package main

import (
	"crm-management/systems/amocrm"
	"fmt"
)

func main() {

	amocrm := amocrm.Amocrm{}
	result := amocrm.Authentificate("mkrivickiy@gmail.com", "470710cae5dedc3a48713b0f07145125")
	fmt.Println(result)
}
