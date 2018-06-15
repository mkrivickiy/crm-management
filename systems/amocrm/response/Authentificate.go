package response

import (
	"crm-management/systems/amocrm/response/authentificate"
)

type Authentificate struct {
	Response   authentificate.Response `json:"response"`
	User       authentificate.User     `json:"user"`
	ServerTime string                  `json:"server_time"`
}
