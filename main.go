package main

import (
	"crm-management/systems/amocrm"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {}

func main() {

	amocrm := amocrm.Amocrm{}
	result := amocrm.Authentificate("mkrivickiy@gmail.com", "470710cae5dedc3a48713b0f07145125")
	fmt.Println(result)
	// router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)
	// http.Handle("/", router)

	// log.Fatal(http.ListenAndServe("prolx.loc:1234", nil))

}
