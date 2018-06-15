package authentificate

import "github.com/mkrivickiy/crm-management/systems/amocrm/authentificate/response/authentificate/response"

type Response struct {
	Auth     bool               `json:"auth"`
	Accounts []response.Account `json:"accounts"`
}
