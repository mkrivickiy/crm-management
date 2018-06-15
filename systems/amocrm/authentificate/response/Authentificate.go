package response

import "github.com/mkrivickiy/crm-management/systems/amocrm/authentificate/response/authentificate"

type Authentificate struct {
	Response   authentificate.Response `json:"response"`
	User       authentificate.User     `json:"user"`
	ServerTime int                     `json:"server_time"`
}
