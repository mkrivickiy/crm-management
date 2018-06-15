package amocrm

import (
	"crm-management/systems/amocrm/response"
	"encoding/json"
	"net/http"
)

type Amocrm struct {
}

func (a *Amocrm) Authentificate(userLogin string, userHash string) *response.Authentificate {
	auth := Auth{Username: userLogin, Userhash: userHash}

	request, _ := http.Post("https://mkrivickiy.amocrm.ru/private/api/auth.php?type=json", "application/json; charset=utf-8", auth.EncodeJson())

	defer request.Body.Close()

	response := &response.Authentificate{}

	_ = json.NewDecoder(request.Body).Decode(response)

	return response

}
