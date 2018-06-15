package authentificate

import (
	"crm-management/systems/amocrm/response/authentificate/response"
)

type Response struct {
	Auth     bool               `json:"auth"`
	Accounts []response.Account `json:"accounts"`
}
